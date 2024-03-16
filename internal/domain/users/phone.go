package users

import (
	"regexp"

	"diplom-backend/internal/common/errors"
)

type Phone string

var phoneRegexp = regexp.MustCompile("^[+]?[(]?[0-9]{3}[)]?[-\\s.]?[0-9]{3}[-\\s.]?[0-9]{4,6}$")

func NewPhoneFromString(s string) (Phone, error) {
	if !phoneRegexp.MatchString(s) {
		return "", errors.NewInvalidInputError("invalid phone format", "phone")
	}

	return Phone(s), nil
}
