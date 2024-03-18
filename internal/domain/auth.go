package domain

import "diplom-backend/internal/domain/validations"

type SignUpRequest struct {
	Name     string
	Phone    string
	Password string
	Type     int16

	Surname *string
	Email   *string
}

func (r SignUpRequest) Validate() error {
	if err := validations.ValidatePhone(r.Phone); err != nil {
		return err
	}

	if r.Email != nil {
		if err := validations.ValidateEmail(*r.Email); err != nil {
			return err
		}
	}

	return nil
}

type SignUpResponse struct {
	Token string
}

type SignInRequest struct {
	Phone    string
	Password string
}

func (r SignInRequest) Validate() error {
	if err := validations.ValidatePhone(r.Phone); err != nil {
		return err
	}
	return nil
}

type SignInResponse struct {
	ID                  int64
	Token               string
	Name                string
	Type                int16
	Surname, Patronymic *string
}
