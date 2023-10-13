package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBParams struct {
	Host     string        `yaml:"host" env-default:"mongodb://localhost:27017" env-required:"true"`
	Port     int           `yaml:"port" env-default:"5432"`
	Timeout  time.Duration `yaml:"timeout" env-default:"10s"`
	User     string        `yaml:"user"`
	Password string        `yaml:"password"`
	DBName   string        `yaml:"dbname" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"15s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Config struct {
	CacheVolume uint8 `yaml:"cache-volume"`
	DBParams    `yaml:"database"`
	HTTPServer  `yaml:"http-server"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH must be set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannnot read config: %s", err)
	}
	return &cfg
}
