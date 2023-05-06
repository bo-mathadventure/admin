package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewAdminMapHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminMap(ctx, db))
	app.Post("/", postAdminMap(ctx, db))
	app.Get("/:id", getAdminMapID(ctx, db))
	app.Put("/:id", putAdminMapID(ctx, db))
	app.Delete("/:id", deleteAdminMapID(ctx, db))
}

type AdminMapResponse struct {
	ID          int      `json:"id"`
	RoomName    string   `json:"roomName" example:"Default Room"`
	MapUrl      string   `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj"`
	Policy      string   `json:"policy" enums:"anonymous,login,permission"`
	ContactPage string   `json:"contactPage" example:"https://mycompany.com/contact-us"`
	Tags        []string `json:"tags" example:"admin,editor"`
	Chat        struct {
		Enable                 bool `json:"enable"`
		EnableUpload           bool `json:"enableUpload"`
		EnableOnlineList       bool `json:"enableOnlineList"`
		EnableDisconnectedList bool `json:"enableDisconnectedList"`
	} `json:"chat"`
	CanReport bool      `json:"canReport"`
	ExpireOn  time.Time `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
	CreatedAt time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
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
		// utils.PERMISSION_MAP_VIEW || utils.PERMISSION_MAP_EDIT
		return handler.HandleSuccess(c)
	}
}

type CreateMap struct {
	RoomName    string   `json:"roomName" example:"Default Room"`
	MapUrl      string   `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj"`
	Policy      string   `json:"policy" enums:"anonymous,login,permission"`
	ContactPage string   `json:"contactPage" example:"https://mycompany.com/contact-us"`
	Tags        []string `json:"tags" example:"admin,editor"`
	Chat        struct {
		Enable                 bool `json:"enable"`
		EnableUpload           bool `json:"enableUpload"`
		EnableOnlineList       bool `json:"enableOnlineList"`
		EnableDisconnectedList bool `json:"enableDisconnectedList"`
	} `json:"chat"`
	CanReport bool      `json:"canReport"`
	ExpireOn  time.Time `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
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
		// utils.PERMISSION_MAP_EDIT
		return handler.HandleSuccess(c)
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
		// utils.PERMISSION_MAP_VIEW || utils.PERMISSION_MAP_EDIT
		return handler.HandleSuccess(c)
	}
}

type UpdateMap struct {
	RoomName    string   `json:"roomName" example:"Default Room"`
	MapUrl      string   `json:"mapUrl" example:"/_/global/thecodingmachine.github.io/workadventure-map-starter-kit/map.tmj"`
	Policy      string   `json:"policy" enums:"anonymous,login,permission"`
	ContactPage string   `json:"contactPage" example:"https://mycompany.com/contact-us"`
	Tags        []string `json:"tags" example:"admin,editor"`
	Chat        struct {
		Enable                 bool `json:"enable"`
		EnableUpload           bool `json:"enableUpload"`
		EnableOnlineList       bool `json:"enableOnlineList"`
		EnableDisconnectedList bool `json:"enableDisconnectedList"`
	} `json:"chat"`
	CanReport bool      `json:"canReport"`
	ExpireOn  time.Time `json:"expireOn" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
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
		// utils.PERMISSION_MAP_EDIT
		return handler.HandleSuccess(c)
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
		// utils.PERMISSION_MAP_EDIT
		return handler.HandleSuccess(c)
	}
}
