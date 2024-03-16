package users

import (
	"time"
	"unicode/utf8"

	"diplom-backend/internal/common/errors"
)

type User struct {
	ID         string
	Name       string
	Surname    string
	Patronymic string
	Age        int16
	Gender     Gender
	Phone      Phone
	Email      Email
	ImageID    string
	CreatedAt  time.Time
	LastOnline *time.Time
	DeletedAt  *time.Time

	PasswordEncrypted string
}

func (u *User) Validate() error {
	if len(u.Name) == 0 || utf8.RuneCountInString(u.Name) > 25 {
		return errors.NewInvalidInputError(errors.StringLengthErrorMessage(0, 40), "name")
	}

	return nil
}
