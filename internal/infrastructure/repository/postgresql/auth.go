package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"diplom-backend/internal/domain"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       type,
		       name,
		       surname,
		       patronymic,
		       password_enc
		from users
		where phone = $1`, phone).Scan(
		&user.ID,
		&user.Type,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.PasswordEncrypted,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &user, nil
}

func (r *AuthRepository) CheckPhoneUnique(ctx context.Context, phone string) error {
	var v int

	err := r.db.QueryRow(ctx, `
		select 1
		from users
		where phone = $1`, phone,
	).Scan(&v)
	if err != nil {
		return parseError(err, "selecting phone")
	}

	return nil
}

func (r *AuthRepository) CheckEmailUnique(ctx context.Context, email string) error {
	var v int

	err := r.db.QueryRow(ctx, `
		select 1
		from users
		where email = $1`, email,
	).Scan(&v)
	if err != nil {
		return parseError(err, "selecting email")
	}

	return nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users(name, type, surname, patronymic, age, gender, image_id, phone, email, last_online, created_at, password_enc)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		user.Name,
		user.Type,
		user.Surname,
		user.Patronymic,
		user.Age,
		user.Gender,
		user.ImageID,
		user.Phone,
		user.Email,
		user.LastOnline,
		user.CreatedAt,
		user.PasswordEncrypted,
	)
	if err != nil {
		return parseError(err, "inserting user")
	}

	return nil
}
