package main

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	_ "github.com/bo-mathadventure/admin/docs"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/handler/admin"
	"github.com/bo-mathadventure/admin/handler/workadventure"
	"github.com/bo-mathadventure/admin/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"time"
)

//	@title						Workadventure Admin Back Office API
//	@version					1.0
//	@description				API documentation for the workdadventure back office written at the Hochschule Bochum
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				GitHub Issues
//	@contact.url				https://github.com/bo-mathadventure/admin
//	@license.name				AGPL 3.0
//	@license.url				https://github.com/teamdigitale/licenses/blob/master/AGPL-3.0-or-later
//	@host						localhost:4664
//	@BasePath					/
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-API-Key
//	@description				JWT user token from login
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

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: config.GetConfig().AppName,
		AppName:      config.GetConfig().AppName,
	})
	app.Use(cors.New())
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
	handler.NewSAMLHandler(authRoute.Group("/saml"), context.Background(), client)

	apiv1 := app.Group("/system", middleware.JWTProtected())
	handler.NewUserHandler(apiv1.Group("/user"), context.Background(), client)

	adminApi := apiv1.Group("/admin", middleware.JWTProtected())
	admin.NewAdminUserHandler(adminApi.Group("/user"), context.Background(), client)
	admin.NewAdminBanHandler(adminApi.Group("/ban"), context.Background(), client)
	admin.NewAdminReportHandler(adminApi.Group("/report"), context.Background(), client)
	admin.NewAdminMapHandler(adminApi.Group("/map"), context.Background(), client)
	admin.NewAdminTextureHandler(adminApi.Group("/texture"), context.Background(), client)
	admin.NewAdminAnnouncementHandler(adminApi.Group("/announcement"), context.Background(), client)
	admin.NewAdminGroupHandler(adminApi.Group("/group"), context.Background(), client)

	waApi := app.Group("/api")
	workadventure.NewRoomHandler(waApi.Group("/room", middleware.AdminAPIProtected()), context.Background(), client)
	workadventure.NewTextureHandler(waApi.Group("/woka"), context.Background(), client)
	workadventure.NewTextureHandler(waApi.Group("/companion"), context.Background(), client)
	workadventure.NewMapHandler(waApi.Group("/map", middleware.AdminAPIProtected()), context.Background(), client)
	workadventure.NewCapabilitiesHandler(waApi.Group("/capabilities"), context.Background(), client)
	workadventure.NewBanHandler(waApi.Group("/ban", middleware.AdminAPIProtected()), context.Background(), client)
	workadventure.NewReportHandler(waApi.Group("/report", middleware.AdminAPIProtected()), context.Background(), client)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 Not Found")
	})

	app.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
