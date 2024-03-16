package users

import (
	"fmt"

	"diplom-backend/internal/common/errors"
)

type Gender int16

const (
	GenderMan Gender = iota + 1
	GenderWoman
)

func NewGender(v int16) (Gender, error) {
	switch Gender(v) {
	case GenderMan:
		return GenderMan, nil
	case GenderWoman:
		return GenderWoman, nil
	default:
		return 0, errors.NewInvalidInputError(fmt.Sprintf("unknown gender %d", v), "gender")
	}
}
