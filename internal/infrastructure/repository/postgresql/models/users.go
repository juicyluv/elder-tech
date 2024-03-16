package models

import (
	"time"

	"diplom-backend/internal/domain/users"
)

type User struct {
	ID          string
	Name        string
	Surname     string
	Patronymic  string
	Age         int16
	Gender      int16
	ImageID     string
	Phone       string
	Email       string
	PasswordEnc string
	CreatedAt   time.Time
	LastOnline  *time.Time
	DeletedAt   *time.Time
}

func UserToDomain(u *User) *users.User {
	return &users.User{
		ID:                u.ID,
		Name:              u.Name,
		Surname:           u.Surname,
		Patronymic:        u.Patronymic,
		Age:               u.Age,
		Gender:            users.Gender(u.Gender),
		Phone:             users.Phone(u.Phone),
		Email:             users.Email(u.Email),
		ImageID:           u.ImageID,
		CreatedAt:         u.CreatedAt,
		LastOnline:        u.LastOnline,
		DeletedAt:         u.DeletedAt,
		PasswordEncrypted: u.PasswordEnc,
	}
}
