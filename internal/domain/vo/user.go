package vo

import (
	"regexp"

	"diplom-backend/internal/common/errors"
)

var (
	phoneRegexp = regexp.MustCompile("^[+]?[(]?[0-9]{3}[)]?[-\\s.]?[0-9]{3}[-\\s.]?[0-9]{4,6}$")
	emailRegexp = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")
)

type Phone string

func NewPhone(p string) (Phone, error) {
	if !phoneRegexp.MatchString(p) {
		return "", errors.NewInvalidInputError("Неправильный формат телефона.", "phone")
	}
	return Phone(p), nil
}

type Email string

func NewEmail(e string) (Email, error) {
	if !emailRegexp.MatchString(e) {
		return "", errors.NewInvalidInputError("Неправильный формат почты.", "email")
	}
	return Email(e), nil
}

type Gender int16
