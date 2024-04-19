package db

import (
	"context"
	"diplom-backend/internal/domain"
	"fmt"
)

func GetImage(ctx context.Context, id int64) (*domain.Image, error) {
	var i domain.Image

	err := db.QueryRow(ctx, `
		select id, filename
		from images
		where id = $1`, id).Scan(
		&i.ID,
		&i.Filename,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting image: %w", err)
	}

	return &i, nil
}
