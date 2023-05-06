package handler

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	saml2 "github.com/russellhaering/gosaml2"
	dsig "github.com/russellhaering/goxmldsig"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
	"time"
)

func NewSAMLHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	cfg := config.GetConfig()
	t := reflect.ValueOf(&cfg).Elem()
	var missingValues []string
	for i := 0; i < t.NumField(); i++ {
		valueField := t.Field(i)
		typeField := t.Type().Field(i)
		tag := typeField.Tag

		if !strings.HasPrefix(tag.Get("env"), "SAMLV2_") {
			continue
		}
		if valueField.IsZero() {
			missingValues = append(missingValues, tag.Get("env"))
		}
	}

	if len(missingValues) > 0 {
		log.WithFields(log.Fields{
			"missing": missingValues,
		}).Warnf("SAMLv2 is not configured or has has missing configuration options - skipping setting up /auth/saml/* routes")
		return
	}

	certStore := dsig.MemoryX509CertificateStore{
		Roots: []*x509.Certificate{},
	}

	certData, err := os.ReadFile(config.GetConfig().SAMLv2RootCert)
	idpCert, err := x509.ParseCertificate(certData)
	if err != nil {
		panic(err)
	}
	certStore.Roots = append(certStore.Roots, idpCert)

	randomKeyStore := dsig.RandomKeyStoreForTest() // fixme sign key for sign authn request

	sp := &saml2.SAMLServiceProvider{
		IdentityProviderSSOURL:      config.GetConfig().SAMLv2SSOURL,
		IdentityProviderIssuer:      config.GetConfig().SAMLv2EntityID,
		ServiceProviderIssuer:       config.GetConfig().SAMLv2Issuer,
		AssertionConsumerServiceURL: fmt.Sprintf("%s/auth/saml/acs", config.GetConfig().BackendURL),
		SignAuthnRequests:           config.GetConfig().SAMLv2SignAuthnRequests,
		AudienceURI:                 config.GetConfig().SAMLv2AudienceURL,
		IDPCertificateStore:         &certStore,
		SPKeyStore:                  randomKeyStore,
	}

	app.Get("/start", getSAMLStart(ctx, db, sp))
	app.Post("/acs", postSAMLacs(ctx, db, sp))
}

// getSAMLStart godoc
//
//	@Summary		Get SAML Auth URL
//	@Description	Starts a new SAML authentication flow. This route is only available when SAML is correctly configured.
//	@Tags			auth,saml
//	@Accept			json
//	@Produce		json
//	@Success		302	{object}	nil
//	@Failure		500	{object}	APIResponse
//	@Router			/auth/saml/start [get]
func getSAMLStart(ctx context.Context, db *ent.Client, sp *saml2.SAMLServiceProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authURL, err := sp.BuildAuthURL("")
		if err != nil {
			return HandleInternalError(c, err)
		}
		return c.Redirect(authURL)
	}
}

type SAMLResponse struct {
	SAMLResponse string `json:"SAMLResponse" xml:"SAMLResponse" form:"SAMLResponse" json:"SAMLResponse"`
	RelayState   string `json:"RelayState" xml:"RelayState" form:"RelayState" json:"RelayState"`
}

// postSAMLacs godoc
//
//	@Summary		SAML Response Callback
//	@Description	Get SAML response of the IDP. This route is only available when SAML is correctly configured.
//	@Tags			auth,saml
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			params	formData	SAMLResponse	true	"-"
//	@Success		302		{object}	nil
//	@Failure		400		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/auth/saml/acs [post]
func postSAMLacs(ctx context.Context, db *ent.Client, sp *saml2.SAMLServiceProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rData := new(SAMLResponse)
		if err := c.BodyParser(rData); err != nil {
			return HandleBodyParseError(c, err)
		}

		assertionInfo, err := sp.RetrieveAssertionInfo(rData.SAMLResponse)
		if err != nil {
			return HandleInternalError(c, err)
		}

		if assertionInfo.WarningInfo.InvalidTime {
			return HandleError(c, "ERR_SAML_INVALIDTIME")
		}

		if assertionInfo.WarningInfo.NotInAudience {
			return HandleError(c, "ERR_SAML_AUDIENCE")
		}

		log.WithFields(log.Fields{
			"NameID":   assertionInfo.NameID,
			"Values":   assertionInfo.Values,
			"Warnings": assertionInfo.WarningInfo,
		}).Info("SAML done")

		samlUID := assertionInfo.NameID                   // fixme what is the correct "uuid" of a saml user?
		samlEMail := assertionInfo.Values.Get("email")    // fixme what is the best way to get the mail
		username := assertionInfo.Values.Get("firstname") // fixme what is the best way to get the firstname

		foundUsers, err := db.User.Query().Where(user.Or(user.EmailEQ(samlEMail), user.SsoIdentifierEQ(samlUID))).All(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		var loginUser *ent.User

		if len(foundUsers) == 0 {
			// we need to register the new user
			newUser, err := db.User.Create().SetEmail(email.Normalize(samlEMail)).SetPassword("").SetUsername(username).SetSsoIdentifier(samlUID).Save(ctx)
			if err != nil || newUser == nil {
				return HandleInternalError(c, err)
			}

			loginUser = newUser

			log.WithFields(log.Fields{
				"userID": loginUser.ID,
			}).Info("saml user registered")
		} else if len(foundUsers) == 1 {
			theUser := foundUsers[0]
			theUserUpdate := theUser.Update()
			if theUser.SsoIdentifier == "" {
				theUserUpdate = theUserUpdate.SetSsoIdentifier(samlUID)
			}
			theUserUpdated, err := theUserUpdate.SetUsername(username).SetLastLogin(time.Now()).Save(ctx)
			if err != nil {
				return HandleInternalError(c, err)
			}
			loginUser = theUserUpdated

			log.WithFields(log.Fields{
				"userID": loginUser.ID,
			}).Info("saml user login")
		} else {
			return HandleError(c, "ERR_SAML_USERCOUNT")
		}

		if loginUser == nil {
			return HandleError(c, "ERR_SAML_NO_USER")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":            loginUser.ID,
			"email":         loginUser.Email,
			"ssoIdentifier": loginUser.SsoIdentifier,
			"exp":           time.Now().Add(time.Hour * 72).Unix(),
		})

		t, err := token.SignedString([]byte(config.GetConfig().WorkadventureSecretKey))
		if err != nil {
			return HandleInternalError(c, err)
		}

		return c.Redirect(fmt.Sprintf("%s?token=%s", config.GetConfig().FrontendURL, t))
	}
}
