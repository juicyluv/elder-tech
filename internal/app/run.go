package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"

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
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "http://localhost:*", "http://127.0.0.1:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Session-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
	r.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	r.Use(middleware.NoCache)

	r.Use(handlers.AuthMiddleware)

	r.Get("/docs/http", handler.DocsFile)
	r.Get("/docs", handler.DocsPage)
	r.Get("/static/*", handler.Static)

	a.httpServer = &stdhttp.Server{
		Addr:    fmt.Sprintf(":" + config.HttpPort()),
		Handler: handlers.HandlerFromMuxWithBaseURL(handler, r, "/api/v1"),
	}

	httpServerCh := make(chan error)
	httpsServerCh := make(chan error)

	go func() {
		httpServerCh <- a.httpServer.ListenAndServe()
	}()

	var httpsServer *http.Server
	if config.HttpsCertPath() != "" && config.HttpsKeyPath() != "" {
		httpsServer = &stdhttp.Server{
			Addr:    fmt.Sprintf(":" + config.HttpsPort()),
			Handler: handlers.HandlerFromMuxWithBaseURL(handler, r, "/api/v1"),
		}

		go func() {
			httpsServerCh <- httpsServer.ListenAndServeTLS(config.HttpsCertPath(), config.HttpsKeyPath())
		}()

		slog.Info(
			"HTTPS server is started",
			slog.String("addr", httpsServer.Addr),
		)
	}

	slog.Info(
		"HTTP server is started",
		slog.String("addr", a.httpServer.Addr),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("Interrupt signal: " + s.String())
	case err = <-httpServerCh:
		slog.Error("HTTP server stop signal: " + err.Error())
	case err = <-httpsServerCh:
		slog.Error("HTTPS server stop signal: " + err.Error())
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownCancel()

	// Shutdown
	err = a.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("failed to shutdown the HTTP server: " + err.Error())
	}
	slog.Info("HTTP server has been shut down successfully")

	if httpsServer != nil {
		err = httpsServer.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("failed to shutdown the HTTPS server: " + err.Error())
		}
		slog.Info("HTTPS server has been shut down successfully")
	}

	return nil
}
