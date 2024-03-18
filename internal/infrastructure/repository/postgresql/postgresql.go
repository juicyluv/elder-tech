package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"diplom-backend/internal/infrastructure/repository"
)

func NewPostgresqlPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("connecting to postgresql: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return pool, nil
}

func parseError(err error, prefix string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		err = repository.ErrNotFound
	}

	if prefix != "" {
		return fmt.Errorf("%s: %w", prefix, err)
	}

	return err
}
