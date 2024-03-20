package usecase

import (
	"context"
	stderrors "errors"
	"fmt"

	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/domain"
	"diplom-backend/internal/infrastructure/repository"
)

type Repository interface {
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

type UserUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("Пользователь не найден.", "user")
		}
		return nil, fmt.Errorf("getting user %d: %w", id, err)
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	user, err := uc.repo.GetUser(ctx, user.ID)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return errors.NewNotFoundError("Пользователь не найден.", "user")
		}
		return fmt.Errorf("getting user %d: %w", user.ID, err)
	}

	err = uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("updating user %d: %w", user.ID, err)
	}

	return nil
}
