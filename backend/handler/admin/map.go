package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/maps"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func NewAdminMapHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	err := handler.Validate.RegisterValidation("policy_validation", func(fl validator.FieldLevel) bool {
		return utils.Contains(MapPolicyNames, fl.Field().String())
	})
	if err != nil {
		log.WithError(err).WithField("validation", "policy_validation").Panic("failed to setup custom validation")
	}

	app.Get("/", getAdminMap(ctx, db))
	app.Post("/", postAdminMap(ctx, db))
	app.Get("/:id", getAdminMapID(ctx, db))
	app.Put("/:id", putAdminMapID(ctx, db))
	app.Delete("/:id", deleteAdminMapID(ctx, db))
}

var defaultMapPolicy = 0

type AdminMapChat struct {
	Enable                 bool `json:"enable" validate:"required,boolean"`
	EnableUpload           bool `json:"enableUpload" validate:"required,boolean"`
	EnableOnlineList       bool `json:"enableOnlineList" validate:"required,boolean"`
	EnableDisconnectedList bool `json:"enableDisconnectedList" validate:"required,boolean"`
}
type AdminMapResponse struct {
	ID          int          `json:"id"`
	RoomName    string       `json:"roomName" example:"Default Room"`
	MapUrl      string       `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj"`
	Policy      string       `json:"policy" enums:"anonymous,login,permission"`
	ContactPage string       `json:"contactPage" example:"https://mycompany.com/contact-us"`
	Tags        []string     `json:"tags" example:"admin,editor"`
	Chat        AdminMapChat `json:"chat"`
	CanReport   bool         `json:"canReport"`
	ExpireOn    string       `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	CreatedAt   string       `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

var MapPolicyNames = []string{
	"anonymous",
	"login",
	"permission",
}

func responseAdminMapResponse(this *ent.Maps) *AdminMapResponse {
	return &AdminMapResponse{
		ID:          this.ID,
		RoomName:    this.RoomName,
		MapUrl:      this.MapUrl,
		Policy:      MapPolicyNames[this.PolicyNumber],
		ContactPage: this.ContactPage,
		Tags:        this.Tags,
		Chat: AdminMapChat{
			Enable:                 this.EnableChat,
			EnableUpload:           this.EnableChatUpload,
			EnableOnlineList:       this.EnableChatOnlineList,
			EnableDisconnectedList: this.EnableChatDisconnectedList,
		},
		CanReport: this.CanReport,
		ExpireOn:  this.ExpireOn.Format(time.RFC3339),
		CreatedAt: this.CreatedAt.Format(time.RFC3339),
	}
}

func responseAdminMapResponses(this []*ent.Maps) []*AdminMapResponse {
	data := make([]*AdminMapResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminMapResponse(e)
	}
	return data
}

// getAdminMap godoc
//
//	@Summary		List maps
//	@Description	Get all maps. Requires permission admin.map.view or admin.map.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminMapResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/map [get]
func getAdminMap(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_MAP_VIEW, utils.PERMISSION_MAP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.Maps.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminMapResponses(allEntires))
	}
}

type CreateMap struct {
	RoomName    string       `json:"roomName" example:"Default Room" validate:"required,min=3,max=32"`
	MapUrl      string       `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj" validate:"required"`
	Policy      string       `json:"policy" enums:"anonymous,login,permission" validate:"required,policy_validation"`
	ContactPage string       `json:"contactPage" example:"https://mycompany.com/contact-us" validate:"required,url"`
	Tags        []string     `json:"tags" example:"admin,editor" validate:"required"`
	Chat        AdminMapChat `json:"chat" validate:"dive"`
	CanReport   bool         `json:"canReport" validate:"required,boolean"`
	ExpireOn    string       `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"required,rfc3339"`
}

// postAdminMap godoc
//
//	@Summary		Create map
//	@Description	Requires permission admin.map.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateMap	true	"-"
//	@Success		200		{object}	AdminMapResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/map [post]
func postAdminMap(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_MAP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		req := new(CreateMap)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		var expireOn *time.Time
		if req.ExpireOn != "" {
			expireOnParsed, err := time.Parse(time.RFC3339, req.ExpireOn)
			if err != nil {
				expireOn = &expireOnParsed
			}
		}

		newCreated, err := db.Maps.Create().
			SetRoomName(req.RoomName).
			SetMapUrl(req.MapUrl).
			SetPolicyNumber(utils.SliceIndex(req.Policy, MapPolicyNames, &defaultMapPolicy)).
			SetContactPage(req.ContactPage).
			SetTags(req.Tags).
			SetCanReport(req.CanReport).
			SetNillableCreatedAt(expireOn).
			SetEnableChat(req.Chat.Enable).
			SetEnableChatUpload(req.Chat.EnableUpload).
			SetEnableChatDisconnectedList(req.Chat.EnableDisconnectedList).
			SetEnableChatOnlineList(req.Chat.EnableOnlineList).
			Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminMapResponse(newCreated))
	}
}

// getAdminMapID godoc
//
//	@Summary		Get specific map
//	@Description	Get map by ID. Requires permission admin.map.view or admin.map.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Map ID"
//	@Success		200	{object}	AdminMapResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/map/{id} [get]
func getAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_MAP_VIEW, utils.PERMISSION_MAP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.Maps.Query().Where(maps.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminMapResponse(foundEntry))
	}
}

type UpdateMap struct {
	RoomName    string       `json:"roomName" example:"Default Room" validate:"required,min=3,max=32"`
	MapUrl      string       `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj" validate:"required"`
	Policy      string       `json:"policy" enums:"anonymous,login,permission" validate:"required,policy_validation"`
	ContactPage string       `json:"contactPage" example:"https://mycompany.com/contact-us" validate:"required,url"`
	Tags        []string     `json:"tags" example:"admin,editor" validate:"required"`
	Chat        AdminMapChat `json:"chat" validate:"dive"`
	CanReport   bool         `json:"canReport" validate:"required,boolean"`
	ExpireOn    string       `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"required,rfc3339"`
}

// putAdminMapID godoc
//
//	@Summary		Update map
//	@Description	Update map by ID. Requires permission admin.map.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateMap	true	"-"
//	@Param			id		path		int			true	"Map ID"
//	@Success		200		{object}	AdminMapResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/map/{id} [put]
func putAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_MAP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(UpdateMap)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.Maps.Query().Where(maps.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		var expireOn *time.Time
		if req.ExpireOn != "" {
			expireOnParsed, err := time.Parse(time.RFC3339, req.ExpireOn)
			if err != nil {
				expireOn = &expireOnParsed
			}
		}

		newFound, err := found.Update().
			SetRoomName(req.RoomName).
			SetMapUrl(req.MapUrl).
			SetPolicyNumber(utils.SliceIndex(req.Policy, MapPolicyNames, &defaultMapPolicy)).
			SetContactPage(req.ContactPage).
			SetTags(req.Tags).
			SetCanReport(req.CanReport).
			SetNillableCreatedAt(expireOn).
			SetEnableChat(req.Chat.Enable).
			SetEnableChatUpload(req.Chat.EnableUpload).
			SetEnableChatDisconnectedList(req.Chat.EnableDisconnectedList).
			SetEnableChatOnlineList(req.Chat.EnableOnlineList).
			Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminMapResponse(newFound))
	}
}

// deleteAdminMapID godoc
//
//	@Summary		Delete map
//	@Description	Delete map by ID. Requires permission admin.map.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Map ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/map/{id} [delete]
func deleteAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_MAP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		err = db.Maps.DeleteOneID(pathID).Exec(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
