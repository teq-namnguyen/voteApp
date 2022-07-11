package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"github.com/namnguyen/backend/codetype"
)

var config *Config

type Config struct {
	Port        string             `envconfig:"PORT"`
	IsDebug     bool               `envconfig:"IS_DEBUG"`
	Stage       codetype.StageType `envconfig:"STAGE"`
	ServiceHost string             `envconfig:"SERVICE_HOST"`
	SercretKey  string             `envconfig:"SECRET_KEY"`
	Postgresql  struct {
		Host           string `envconfig:"DB_HOST"`
		Port           string `envconfig:"DB_PORT"`
		User           string `envconfig:"DB_USER"`
		Pass           string `envconfig:"DB_PASS"`
		DBName         string `envconfig:"DB_NAME"`
		DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
		CountRetryTx   int    `envconfig:"DB_TX_RETRY_COUNT"`
	}
}

func init() {
	config = &Config{}

	_ = godotenv.Load()

	err := envconfig.Process("", config)
	if err != nil {
		err = errors.Wrap(err, "Failed to decode config env")
		log.Fatal(err.Error())
	}

	config.Stage.UpCase()
}

func GetConfig() *Config {
	return config
}
