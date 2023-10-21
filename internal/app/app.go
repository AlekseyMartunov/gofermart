package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"AlekseyMartunov/internal/accrualservice/accural"
	"AlekseyMartunov/internal/accrualservice/collector"
	"AlekseyMartunov/internal/accrualservice/requestcontroller"
	"AlekseyMartunov/internal/adapters/db/migration"
	postgres2 "AlekseyMartunov/internal/adapters/db/orders/postgres"
	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"AlekseyMartunov/internal/adapters/http/loginhandlers"
	"AlekseyMartunov/internal/adapters/http/orderhandlers"
	"AlekseyMartunov/internal/adapters/http/router"
	"AlekseyMartunov/internal/adapters/http/userhandlers"
	"AlekseyMartunov/internal/config"
	"AlekseyMartunov/internal/logger"
	"AlekseyMartunov/internal/middleware/auth"
	"AlekseyMartunov/internal/middleware/login"
	"AlekseyMartunov/internal/orders"
	"AlekseyMartunov/internal/tokenmanager"
	"AlekseyMartunov/internal/users"
	"AlekseyMartunov/internal/utils/hashencoder"
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

	userRepo := postgres.NewUserStorage(conn, logger)
	orderRepo := postgres2.NewOrderStorage(conn, logger)

	hash := hashencoder.New()

	userService := users.NewUserService(userRepo, hash)
	orderService := orders.NewOrderService(orderRepo)

	tokenController := tokenmanager.New(time.Hour*10, []byte("Secret key"))

	auth := auth.New(userService, tokenController, logger)
	logMiddleware := login.NewLoggerMiddleware(logger)

	orderHandler := orderhandlers.New(logger, userService, orderService)
	userHandler := userhandlers.New(logger, userService, orderService)
	loginHandler := loginhandlers.NewLoginHandler(logger, userService, tokenController)

	router := router.NewRouter(userHandler, orderHandler, loginHandler, auth, logMiddleware)

	collector := collector.NewCollector(conn, logger)
	req := requestcontroller.New(cfg.Accrual(), "/api/orders/", logger)
	acc := accural.NewAccrual(logger, orderService)

	numberChan := collector.Run(ctx, time.Second*3)
	orderChan := req.Run(ctx, numberChan)
	acc.Run(ctx, orderChan, time.Second*3)

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
	//dsn := "postgres://admin:1234@localhost:5432/test"
	dsn := cfg.DSN()
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

func connection(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	//dsn := "postgres://admin:1234@localhost:5432/test"
	dsn := cfg.DSN()
	//conn, err := pgx.Connect(ctx, dsn)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("error with connect to db: %w", err)
	}
	return conn, nil
}
