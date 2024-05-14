package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	"diplom-backend/internal/common/config"
	"diplom-backend/internal/db"
	"diplom-backend/internal/filesystem"
	"diplom-backend/internal/handlers"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (a *App) Run(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("creating pgxpool: %w", err)
	}
	defer pool.Close()

	if err = pool.Ping(ctx); err != nil {
		return fmt.Errorf("pinging database: %w", err)
	}
	slog.Info("Connected to database")

	m, err := migrate.New("file://migrations", config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("creating migration: %w", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("applying migrations: %w", err)
	}

	db.Init(pool)

	imageFileSys, err := filesystem.New("images")
	if err != nil {
		return fmt.Errorf("creating images fs")
	}

	handler := handlers.NewHandler(
		imageFileSys,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(handlers.AuthMiddleware)

	// CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://127.0.0.1:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	r.Get("/docs/http", handler.DocsFile)
	r.Get("/docs", handler.DocsPage)
	r.Get("/static/*", handler.Static)

	a.httpServer = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort()),
		Handler: handlers.HandlerFromMuxWithBaseURL(handler, r, "/api/v1"),
	}

	httpServerCh := make(chan error)
	go func() {
		httpServerCh <- a.httpServer.ListenAndServe()
	}()

	slog.Info(
		"Server is started",
		slog.String("addr", a.httpServer.Addr),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("Interrupt signal: " + s.String())
	case err = <-httpServerCh:
		slog.Error("Server stop signal: " + err.Error())
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownCancel()

	// Shutdown
	err = a.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("failed to shutdown the server: " + err.Error())
	}
	slog.Info("Server has been shut down successfully")

	return nil
}
