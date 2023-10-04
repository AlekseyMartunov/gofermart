package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"

	"AlekseyMartunov/internal/adapters/db/migration"
	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"AlekseyMartunov/internal/adapters/http/handlers"
	"AlekseyMartunov/internal/adapters/http/router"
	"AlekseyMartunov/internal/users"
	"AlekseyMartunov/internal/utils/config"
	"AlekseyMartunov/internal/utils/hashencoder"
	"AlekseyMartunov/internal/utils/logger"
)

func StartApp(ctx context.Context) error {

	cfg := config.New()
	cfg.ParseFlags()

	if err := runMigrations(cfg); err != nil {
		return err
	}

	logger, err := logger.New()
	if err != nil {
		return err
	}
	defer logger.Sync()

	conn, err := connection(ctx, cfg)
	if err != nil {
		return err
	}

	repo := postgres.NewUserStorage(conn, logger)
	hash := hashencoder.New()
	userService := users.NewUserService(repo, hash)

	handler := handlers.New(logger, userService)
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

func runMigrations(cfg *config.Config) error {
	dsn := "postgres://admin:1234@localhost:5432/test"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("error with connect to db: %w", err)
	}
	defer db.Close()

	err = migration.StartMigrations(db)
	if err != nil {
		return fmt.Errorf("migration error: %w", err)
	}

	return nil
}

func connection(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	dsn := "postgres://admin:1234@localhost:5432/test"
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("error with connect to db: %w", err)
	}
	return conn, nil
}
