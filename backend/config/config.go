package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var cfg Config

type Config struct {
	Port        int    `env:"PORT" envDefault:"4664"`
	AppName     string `env:"APP_NAME" envDefault:"WorkAdventure Back Office"`
	FrontendURL string `env:"FRONTEND_URL" envDefault:"http://localhost"`
	BackendURL  string `env:"BACKEND_URL" envDefault:"http://localhost"`

	DatabaseType     string `env:"DB_TYPE" envDefault:"mysql"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     int    `env:"DB_PORT" envDefault:"3306"`
	DatabaseName     string `env:"DB_NAME"`
	DatabaseUsername string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`

	JWTSecret string `env:"JWT_SECRET"`

	WorkadventureURL            string `env:"WORKADVENTURE_URL"`
	WorkadventureURLProtocol    string `env:"WORKADVENTURE_URL_PROTOCOL" envDefault:"https"`
	WorkadventureAdminAPISecret string `env:"WORKADVENTURE_ADMIN_API_SECRET"`
	WorkadventureStartRoomURL   string `env:"WORKADVENTURE_START_ROOM_URL"`

	MapStorageURL      string `env:"MAP_STORAGE_URL"`
	MapStorageUser     string `env:"MAP_STORAGE_USER"`
	MapStoragePassword string `env:"MAP_STORAGE_PASSWORD"`
}

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Warn("no .env file found - skipping")
	}

	cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return cfg
}

func GetDBUri(tcp bool) string {
	template := "%s:%s@%s:%d/%s?parseTime=True"
	if tcp {
		template = "%s:%s@tcp(%s:%d)/%s?parseTime=True"
	}
	return fmt.Sprintf(template,
		cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
}
