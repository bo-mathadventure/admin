package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var cfg Config

type Config struct {
	Port        int    `env:"PORT" envDefault:"4664" json:"-"`
	AppName     string `env:"APP_NAME" envDefault:"WorkAdventure Back Office" json:"appName"`
	FrontendURL string `env:"FRONTEND_URL" envDefault:"http://localhost" json:"-"`
	BackendURL  string `env:"BACKEND_URL" envDefault:"http://localhost" json:"-"`

	DatabaseType     string `env:"DB_TYPE" envDefault:"mysql" json:"-"`
	DatabaseHost     string `env:"DB_HOST" json:"-"`
	DatabasePort     int    `env:"DB_PORT" envDefault:"3306" json:"-"`
	DatabaseName     string `env:"DB_NAME" json:"-"`
	DatabaseUsername string `env:"DB_USER" json:"-"`
	DatabasePassword string `env:"DB_PASSWORD" json:"-"`

	EnableRegistration bool `env:"ENABLE_REGISTRATION" envDefault:"true" json:"enableRegistration"`

	WorkadventureURL            string `env:"WORKADVENTURE_URL" json:"workadventureURL"`
	WorkadventureURLProtocol    string `env:"WORKADVENTURE_URL_PROTOCOL" envDefault:"https" json:"workadventureURLProtocol"`
	WorkadventureAdminAPISecret string `env:"WORKADVENTURE_ADMIN_API_SECRET" json:"-"`
	WorkadventureStartRoomURL   string `env:"WORKADVENTURE_START_ROOM_URL" json:"-"`
	WorkadventureSecretKey      string `env:"WORKADVENTURE_SECRET_KEY" json:"-"`

	MapStorageURL      string `env:"MAP_STORAGE_URL" json:"mapStorageURL"`
	MapStorageUser     string `env:"MAP_STORAGE_USER" json:"-"`
	MapStoragePassword string `env:"MAP_STORAGE_PASSWORD" json:"-"`

	SAMLv2RootCert          string `env:"SAMLV2_ROOT_CERT" json:"-"`
	SAMLv2SSOURL            string `env:"SAMLV2_SSO_URL" json:"-"`
	SAMLv2EntityID          string `env:"SAMLV2_ENTITYID" json:"-"`
	SAMLv2Issuer            string `env:"SAMLV2_ISSUER_URL" json:"-"`
	SAMLv2AudienceURL       string `env:"SAMLV2_AUDIENCE_URL" json:"-"`
	SAMLv2SignAuthnRequests bool   `env:"SAMLV2_SIGN_AUTH_REQUESTS" envDefault:"true" json:"-"`
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
