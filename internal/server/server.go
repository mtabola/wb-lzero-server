package server

import (
	"log"
	"net/http"
	"server/config"

	"server/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MustLoadServer(cfg *config.HTTPServer) {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/orders", func(r chi.Router) {
		r.Get("/", handlers.GetAllOrders)
		r.Post("/", handlers.SaveOrder)
	})

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server error: ", err)
	}
}
