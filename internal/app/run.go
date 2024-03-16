package app

import (
	"context"
	"fmt"
	"log"
	stdhttp "net/http"

	"diplom-backend/internal/common/config"
	"diplom-backend/internal/handlers/http"
	"diplom-backend/internal/infrastructure/repository/postgresql"
	"diplom-backend/internal/usecase/users"
)

func (a *App) Run(ctx context.Context) error {
	db, err := postgresql.NewPostgresqlPool(ctx, config.DatabaseURL())
	if err != nil {
		return fmt.Errorf("connecting to postgresql: %w", err)
	}
	defer db.Close() // TODO: mb in shutdown

	userRepository := postgresql.NewUserRepository(db)
	// TODO: filesystem

	// tODO: usecases
	userUseCase := users.NewUseCase(userRepository)
	// TODO: handlers
	handler := http.NewHandler(userUseCase)
	// TODO: http server

	a.httpServer = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort()),
		Handler: http.HandlerFromMux(handler, nil),
	}

	log.Println("running server")
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	//select {
	//case <-ctx.Done():
	//	// TODO: case done chan
	//	// TODO: case err chan
	//}

	// TODO: run server
	// TODO: shutdown

	return nil
}
