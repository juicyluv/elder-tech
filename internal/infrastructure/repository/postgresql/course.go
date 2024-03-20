package postgresql

import (
	"context"

	"diplom-backend/internal/domain"
)

func (r *Repository) CreateCourse(ctx context.Context, course *domain.Course) (int32, error) {
	return 0, nil
}
