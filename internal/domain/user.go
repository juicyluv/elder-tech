package domain

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"diplom-backend/internal/common/errors"
)

var (
	phoneRegexp = regexp.MustCompile("^[+]?[(]?[0-9]{3}[)]?[-\\s.]?[0-9]{3}[-\\s.]?[0-9]{4,6}$")
	emailRegexp = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")
)

type User struct {
	ID                int64
	Name              string
	Phone             string
	PasswordEncrypted string
	CreatedAt         time.Time

	Surname    *string
	Patronymic *string
	Age        *int16
	Gender     *int16
	Email      *string
	ImageID    *int64
	LastOnline *time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.NewInvalidInputError("Имя не может быть пустым.", "name")
	}
	if utf8.RuneCountInString(u.Name) > 30 {
		return errors.NewInvalidInputError("Имя слишком длинное.", "name")
	}

	if !phoneRegexp.MatchString(u.Phone) {
		return errors.NewInvalidInputError("Неправильный формат телефона.", "phone")
	}

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
