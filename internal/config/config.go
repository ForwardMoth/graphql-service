package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	DBConfig
	ServerPort   string `env:"HTTP_PORT" env-default:"8080"`
	InMemoryMode bool   `env:"IN_MEMORY_MODE" env-default:"false"`
}

type DBConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal("Config file doesn't exist, .env")
	}

	var config Config
	if err := cleanenv.ReadConfig(".env", &config); err != nil {
		log.Fatal("Can't read config file: .env")
	}

	return &config
}
