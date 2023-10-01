package app

import (
	"context"
	"fmt"
	"net/http"

	"AlekseyMartunov/internal/adapters/http/handlers"
	"AlekseyMartunov/internal/adapters/http/router"
	"AlekseyMartunov/internal/logger"
)

func StartApp(ctx context.Context) error {

	logger, err := logger.New()
	if err != nil {
		return fmt.Errorf("creating logger error: %w", err)
	}
	defer logger.Sync()

	handler := handlers.New()
	router := router.NewRouter(handler)

	s := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router.Route(),
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve error: %w", err)
	}

	return nil
}
