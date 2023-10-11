package main

import (
	"os"
	"server/config"
	"server/internal/server"
)

func main() {
	os.Setenv("CONFIG_PATH", "../config/config.yaml")
	cfg := config.MustLoadConfig()

	server.MustLoadServer(&cfg.HTTPServer)
}
