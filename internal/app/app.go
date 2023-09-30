package app

import (
	"context"
	"log"
	"net/http"

	"AlekseyMartunov/internal/adapters/http/handlers"
	"AlekseyMartunov/internal/adapters/http/router"
)

func StartApp(ctx context.Context) {

	handler := handlers.New()
	router := router.NewRouter(handler)

	s := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router.Route(),
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
