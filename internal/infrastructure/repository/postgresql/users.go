package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"diplom-backend/internal/domain/users"
	"diplom-backend/internal/infrastructure/repository/postgresql/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*users.User, error) {
	var u models.User

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
		       last_online,
		       deleted_at
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
		&u.DeletedAt,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return models.UserToDomain(&u), nil
}
