package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"diplom-backend/internal/domain/vo"
)

type User struct {
	ID                int64
	Name              string
	Phone             vo.Phone
	Type              int16
	PasswordEncrypted string
	CreatedAt         time.Time

	Surname    *string
	Patronymic *string
	Age        *int16
	Gender     *int16
	Email      *vo.Email
	ImageID    *int64
	LastOnline *time.Time
}

func (u *User) Validate() error {
	return nil
}

func (u *User) EncryptPassword() error {
	data, err := bcrypt.GenerateFromPassword([]byte(u.PasswordEncrypted), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("encrypting password: %w", err)
	}

	u.PasswordEncrypted = string(data)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordEncrypted), []byte(password)) == nil
}
