package db

import (
	"context"
	"fmt"

	"diplom-backend/internal/domain"
)

func GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user domain.User

	err := db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic,
		       password_enc
		from users
		where phone = $1`, phone).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.PasswordEncrypted,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting user: %w", err)
	}

	return &user, nil
}

func GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error) {
	ac := domain.AuthContext{ID: id}

	err := db.QueryRow(ctx, `
		select id, name
		from users
		where id=$1`, id,
	).Scan(&ac.ID, &ac.Name)
	if err != nil {
		return nil, fmt.Errorf("selecting email: %w", err)
	}

	return &ac, nil
}
