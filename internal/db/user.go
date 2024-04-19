package db

import (
	"context"
	"fmt"
	"strings"

	"diplom-backend/internal/domain"
)

func GetUser(ctx context.Context, id int64) (*domain.User, error) {
	var u domain.User

	err := db.QueryRow(ctx, `
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
		return nil, fmt.Errorf("selecting user: %w", err)
	}

	return &u, nil
}

func UpdateUser(ctx context.Context, user *domain.User) error {
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

	_, err := db.Exec(ctx, fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id=$1`, query),
		args...,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}

	return nil
}

func CheckPhoneUnique(ctx context.Context, phone string) error {
	var v int

	err := db.QueryRow(ctx, `
		select 1
		from users
		where phone = $1`, phone,
	).Scan(&v)
	if err != nil {
		return fmt.Errorf("selecting phone: %w", err)
	}

	return nil
}

func CheckEmailUnique(ctx context.Context, email string) error {
	var v int

	err := db.QueryRow(ctx, `
		select 1
		from users
		where email = $1`, email,
	).Scan(&v)
	if err != nil {
		return fmt.Errorf("selecting email: %w", err)
	}

	return nil
}

func CreateUser(ctx context.Context, user *domain.User) error {
	_, err := db.Exec(ctx, `
		INSERT INTO users(name, surname, patronymic, age, gender, image_id, phone, email, last_online, created_at, password_enc)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		user.Name,
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
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}

func UpdateUserImage(ctx context.Context, userID int64, imageFilename string, oldImageID *int64) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("beginning tx: %w", err)
	}
	defer tx.Rollback(context.Background())

	var imageID int64
	err = tx.QueryRow(ctx, "INSERT INTO images(filename) VALUES($1) RETURNING id", imageFilename).Scan(&imageID)
	if err != nil {
		return fmt.Errorf("inserting image: %w", err)
	}

	_, err = tx.Exec(ctx, "UPDATE users SET image_id=$2 WHERE id=$1", userID, imageID)
	if err != nil {
		return fmt.Errorf("updating user image id: %w", err)
	}

	if oldImageID != nil {
		_, err = tx.Exec(ctx, "DELETE FROM images WHERE id=$1", oldImageID)
		if err != nil {
			return fmt.Errorf("deleting old image: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("committing tx: %w", err)
	}

	return nil
}
