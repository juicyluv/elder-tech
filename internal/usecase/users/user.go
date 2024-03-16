package users

import (
	"context"
	stderrors "errors"
	"fmt"

	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/domain/users"
	"diplom-backend/internal/infrastructure/repository"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*users.User, error)
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) GetUser(ctx context.Context, id string) (*users.User, error) {
	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("user not found", "user")
		}

		return nil, fmt.Errorf("getting user %d: %w", id, err)
	}

	return user, nil
}
