package main

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/bo-mathadventure/admin/config"
	_ "github.com/bo-mathadventure/admin/docs"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/token"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/handler/admin"
	"github.com/bo-mathadventure/admin/handler/workadventure"
	"github.com/bo-mathadventure/admin/mailer"
	"github.com/bo-mathadventure/admin/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// @title						Workadventure Admin Back Office API
// @version						1.0
// @description					API documentation for the workdadventure back office written at the Hochschule Bochum
// @termsOfService				http://swagger.io/terms/
// @contact.name				GitHub Issues
// @contact.url					https://github.com/bo-mathadventure/admin
// @license.name				AGPL 3.0
// @license.url					https://github.com/teamdigitale/licenses/blob/master/AGPL-3.0-or-later
// @host						localhost:4664
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-API-Key
// @description					JWT user token from login
func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	if err := config.Init(); err != nil {
		log.WithError(err).Panic("failed to parse init")
	}

	client, err := ent.Open(config.GetConfig().DatabaseType, config.GetDBUri(true))
	if err != nil {
		log.WithFields(log.Fields{
			"dbType": config.GetConfig().DatabaseType,
			"dbPort": config.GetConfig().DatabasePort,
			"dbName": config.GetConfig().DatabaseName,
			"dbUser": config.GetConfig().DatabaseUsername,
		}).WithError(err).Panic("failed opening connection to database")
	}
	defer client.Close()

	err = handler.Validate.RegisterValidation("rfc3339", func(fl validator.FieldLevel) bool {
		_, parseError := time.Parse(time.RFC3339, fl.Field().String())
		return parseError == nil
	})
	if err != nil {
		log.WithError(err).WithField("validation", "rfc3339").Panic("failed to setup custom validation")
	}

	mailer.Init(config.GetConfig())

	var schedulerMutex sync.Mutex
	c := cron.New()
	_, _ = c.AddFunc("* * * * *", func() {
		if !schedulerMutex.TryLock() {
			log.Warn("token cron still locked")
			return
		}
		defer schedulerMutex.Unlock()
		unsend, _ := client.Token.Query().Where(token.Send(false)).WithUser().All(context.Background())

		for _, unsendToken := range unsend {
			userLang := "de"
			log.WithFields(log.Fields{
				"template": unsendToken.Action,
				"userID":   unsendToken.Edges.User.ID,
				"lang":     userLang,
			}).Info("sending email notification")

			merged := make(map[string]interface{})
			for k, v := range unsendToken.Data {
				merged[k] = v
			}
			merged["token"] = unsendToken.Token
			merged["email"] = unsendToken.Edges.User.Email
			merged["username"] = unsendToken.Edges.User.Username
			merged["appName"] = config.GetConfig().AppName
			merged["frontendURL"] = config.GetConfig().FrontendURL
			merged["backendURL"] = config.GetConfig().BackendURL

			err = mailer.Send(filepath.Join("template", userLang, fmt.Sprintf("%s.gohtml", unsendToken.Action)), unsendToken.Edges.User.Email, merged)
			if err != nil {
				log.WithError(err).Error("failed to send email notification")
				continue
			}
			_, err := unsendToken.Update().SetSend(true).Save(context.Background())
			if err != nil {
				log.WithError(err).Error("failed to update token")
				continue
			}

		}
	})
	c.Start()
	defer c.Stop()

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: config.GetConfig().AppName,
		AppName:      config.GetConfig().AppName,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(requestid.New())
	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Static("/public", "./public")

	app.Get("/swagger/*", swagger.HandlerDefault)

	authRoute := app.Group("/auth")
	authRoute.Post("/login", handler.Login(context.Background(), client))
	authRoute.Post("/register", handler.Register(context.Background(), client))
	authRoute.Post("/resendConfirmation", handler.ResendConfirmEmail(context.Background(), client))
	authRoute.Post("/token", handler.Token(context.Background(), client))
	handler.NewSAMLHandler(context.Background(), authRoute.Group("/saml"), client)

	apiv1 := app.Group("/system", middleware.JWTProtected())
	handler.NewUserHandler(context.Background(), apiv1.Group("/user"), client)

	adminAPI := apiv1.Group("/admin", middleware.JWTProtected())
	admin.NewAdminUserHandler(context.Background(), adminAPI.Group("/user"), client)
	admin.NewAdminBanHandler(context.Background(), adminAPI.Group("/ban"), client)
	admin.NewAdminReportHandler(context.Background(), adminAPI.Group("/report"), client)
	admin.NewAdminMapHandler(context.Background(), adminAPI.Group("/map"), client)
	admin.NewAdminTextureHandler(context.Background(), adminAPI.Group("/texture"), client)
	admin.NewAdminAnnouncementHandler(context.Background(), adminAPI.Group("/announcement"), client)
	admin.NewAdminGroupHandler(context.Background(), adminAPI.Group("/group"), client)

	waAPI := app.Group("/api")
	workadventure.NewRoomHandler(context.Background(), waAPI.Group("/room", middleware.AdminAPIProtected()), client)
	workadventure.NewTextureHandler(context.Background(), waAPI.Group("/woka"), client)
	workadventure.NewTextureHandler(context.Background(), waAPI.Group("/companion"), client)
	workadventure.NewMapHandler(context.Background(), waAPI.Group("/map", middleware.AdminAPIProtected()), client)
	workadventure.NewCapabilitiesHandler(context.Background(), waAPI.Group("/capabilities"), client)
	workadventure.NewBanHandler(context.Background(), waAPI.Group("/ban", middleware.AdminAPIProtected()), client)
	workadventure.NewReportHandler(context.Background(), waAPI.Group("/report", middleware.AdminAPIProtected()), client)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 Not Found")
	})

	app.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
