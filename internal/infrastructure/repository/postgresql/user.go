package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"diplom-backend/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	var u domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic,
		       age,
		       gender,
		       image_id,
		       phone,
		       email,
		       created_at,
		       last_online
		from users
		where id = $1`, id).Scan(
		&u.ID,
		&u.Name,
		&u.Surname,
		&u.Patronymic,
		&u.Age,
		&u.Gender,
		&u.ImageID,
		&u.Phone,
		&u.Email,
		&u.CreatedAt,
		&u.LastOnline,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &u, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	args := []any{user.ID}
	var fields []string
	argID := 2

	fields = append(fields, fmt.Sprintf("name=$%d", argID))
	argID++
	args = append(args, user.Name)

	fields = append(fields, fmt.Sprintf("phone=$%d", argID))
	argID++
	args = append(args, user.Phone)

	if user.Age != nil {
		fields = append(fields, fmt.Sprintf("age=$%d", argID))
		argID++
		args = append(args, *user.Age)
	}

	if user.Surname != nil {
		fields = append(fields, fmt.Sprintf("surname=$%d", argID))
		argID++
		args = append(args, *user.Surname)
	}

	if user.Patronymic != nil {
		fields = append(fields, fmt.Sprintf("patronymic=$%d", argID))
		argID++
		args = append(args, *user.Patronymic)
	}

	if user.Gender != nil {
		fields = append(fields, fmt.Sprintf("gender=$%d", argID))
		argID++
		args = append(args, *user.Gender)
	}

	if user.Email != nil {
		fields = append(fields, fmt.Sprintf("email=$%d", argID))
		argID++
		args = append(args, *user.Email)
	}

	query := strings.Join(fields, ",")

	_, err := r.db.Exec(ctx, fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id=$1`, query),
		args...,
	)
	if err != nil {
		return parseError(err, "updating user")
	}

	return nil
}
