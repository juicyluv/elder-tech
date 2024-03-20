package app

import (
	"context"
	"fmt"
	"log/slog"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"diplom-backend/internal/common/config"
	"diplom-backend/internal/handlers/http"
	"diplom-backend/internal/infrastructure/repository/postgresql"
	"diplom-backend/internal/usecase"
)

func (a *App) Run(ctx context.Context) error {
	db, err := postgresql.NewPostgresqlPool(ctx, config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("connecting to postgresql: %w", err)
	}

	slog.Info("Connected to database")

	userRepository := postgresql.NewUserRepository(db)
	userUseCase := usecase.NewUseCase(userRepository)

	authRepository := postgresql.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepository)

	handler := http.NewHandler(
		userUseCase,
		authUseCase,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use()

	a.httpServer = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort()),
		Handler: http.HandlerFromMuxWithBaseURL(handler, r, "/api/v1"),
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
	db.Close()
	slog.Info("Server has been shut down successfully")

	return nil
}
