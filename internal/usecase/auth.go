package usecase

import (
	"context"
	stderrors "errors"
	"fmt"
	"time"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/domain"
	"diplom-backend/internal/infrastructure/repository"
)

type AuthRepository interface {
	GetUserByPhone(ctx context.Context, phone string) (*domain.User, error)
	CheckPhoneUnique(ctx context.Context, phone string) error
	CheckEmailUnique(ctx context.Context, email string) error
	CreateUser(ctx context.Context, user *domain.User) error
	GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error)
}

type AuthUseCase struct {
	authRepo AuthRepository
}

func NewAuthUseCase(authRepo AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo}
}

func (uc *AuthUseCase) SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error) {
	// todo: validate

	user, err := uc.authRepo.GetUserByPhone(ctx, req.Phone)
	if err != nil && !stderrors.Is(err, repository.ErrNotFound) {
		return nil, fmt.Errorf("getting user by phone: %w", err)
	}

	if stderrors.Is(err, repository.ErrNotFound) || !user.ComparePassword(req.Password) {
		return nil, errors.NewAuthError("Неправильный телефон или пароль.", "credentials")
	}

	token, err := auth.GenerateJWT(auth.Claims{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("generating token: %w", err)
	}

	return &domain.SignInResponse{
		Token:      token,
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
	}, nil
}

func (uc *AuthUseCase) SignUp(ctx context.Context, req *domain.SignUpRequest) error {
	// todo: validate

	err := uc.authRepo.CheckPhoneUnique(ctx, req.Phone)
	if err != nil && !stderrors.Is(err, repository.ErrNotFound) {
		return fmt.Errorf("checking phone unique: %w", err)
	}
	if err == nil {
		return errors.NewInvalidInputError("Телефон уже используется.", "phone")
	}

	if req.Email != nil {
		err = uc.authRepo.CheckEmailUnique(ctx, *req.Email)
		if err != nil && !stderrors.Is(err, repository.ErrNotFound) {
			return fmt.Errorf("checking email unique: %w", err)
		}
		if err == nil {
			return errors.NewInvalidInputError("Почта уже используется.", "email")
		}
	}

	user := domain.User{
		Name:              req.Name,
		Phone:             req.Phone,
		PasswordEncrypted: req.Password,
		CreatedAt:         time.Now(),
		Surname:           req.Surname,
		Email:             req.Email,
	}
	if err = user.EncryptPassword(); err != nil {
		return fmt.Errorf("encrypting password: %w", err)
	}

	if err = uc.authRepo.CreateUser(ctx, &user); err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}

func (uc *AuthUseCase) GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error) {
	authContext, err := uc.authRepo.GetAuthContextByUserID(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewAuthError("Пользователь не авторизован.", "token")
		}
		return nil, fmt.Errorf("getting user %d: %w", id, err)
	}

	return authContext, nil
}
