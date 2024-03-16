package users

import (
	"regexp"

	"diplom-backend/internal/common/errors"
)

type Email string

var emailRegexp = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")

func NewEmail(v string) (Email, error) {
	if !emailRegexp.MatchString(v) {
		return "", errors.NewInvalidInputError("invalid email format", "email")
	}

	return Email(v), nil
}
