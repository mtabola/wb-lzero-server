package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBParams struct {
	Host     string        `yaml:"host" env-default:"mongodb://localhost:27017" env-required:"true"`
	Port     string        `yaml:"port" env-default:"5432"`
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

type STANSubscriber struct {
	Server    string `yaml:"server" env-default:"nats://127.0.0.1:4222" env-required:"true"`
	Cluster   string `yaml:"cluster" env-default:"wb-practice" env-required:"true"`
	Channel   string `yaml:"channel" env-default:"lzero"`
	Client    string `yaml:"client" env-default:"service_listener"`
	TimeDelta string `yaml:"time_delta"`
}

type Config struct {
	DBParams       `yaml:"database"`
	HTTPServer     `yaml:"http-server"`
	STANSubscriber `yaml:"stan-sub"`
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
