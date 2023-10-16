package main

import (
	"log"
	"os"
	"server/config"
	"server/internal/cache"
	"server/internal/db"
	"server/internal/handlers"
	"server/internal/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	os.Setenv("CONFIG_PATH", "../config/config.yaml")
	cfg := config.MustLoadConfig()

	db, err := db.New(cfg.DBParams)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	validate := validator.New()

	cache, err := cache.New(cfg.CacheVolume, db)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	handler := handlers.New(cache, validate)

	server.MustLoadServer(&cfg.HTTPServer, handler)
}
