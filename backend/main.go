package main

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/handler/workadventure"
	"github.com/bo-mathadventure/admin/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	log "github.com/sirupsen/logrus"
)

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

	authRoute := app.Group("/auth")
	authRoute.Post("/login", handler.Login(context.Background(), client))
	authRoute.Post("/register", handler.Register(context.Background(), client))

	apiv1 := app.Group("/system", middleware.JWTProtected())
	handler.NewUserHandler(apiv1.Group("/user"), context.Background(), client)

	//waApi := app.Group("/api", middleware.AdminAPIProtected())
	waApi := app.Group("/api")
	workadventure.NewRoomHandler(waApi.Group("/room"), context.Background(), client)
	workadventure.NewTextureHandler(waApi.Group("/woka"), context.Background(), client)
	workadventure.NewTextureHandler(waApi.Group("/companion"), context.Background(), client)
	workadventure.NewMapHandler(waApi.Group("/"), context.Background(), client)
	workadventure.NewCapabilitiesHandler(waApi.Group("/"), context.Background(), client)
	workadventure.NewBanHandler(waApi.Group("/"), context.Background(), client)
	workadventure.NewReportHandler(waApi.Group("/"), context.Background(), client)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 Not Found")
	})

	app.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
