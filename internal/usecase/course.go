package usecase

import (
	"context"
	stderrors "errors"
	"fmt"

	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/domain"
	"diplom-backend/internal/infrastructure/repository"
)

type CourseRepository interface {
	GetUser(ctx context.Context, id int64) (*domain.User, error)

	GetCourse(ctx context.Context, id int32) (*domain.Course, error)
	GetAuthorCourses(ctx context.Context, userID int64) ([]domain.Course, error)
	GetUserCourses(ctx context.Context, userID int64) ([]domain.Course, error)
	CreateCourse(ctx context.Context, course *domain.Course) (int32, error)
	UpdateCourse(ctx context.Context, course *domain.Course) error
	DeleteCourse(ctx context.Context, id int32) error

	GetCourseMembers(ctx context.Context, id int32) ([]domain.User, error)
	AddCourseMember(ctx context.Context, courseID int32, userID int64) error
}

type CourseUseCase struct {
	courseRepository CourseRepository
}

func NewCourseUseCase(courseRepository CourseRepository) *CourseUseCase {
	return &CourseUseCase{courseRepository: courseRepository}
}

func (uc *CourseUseCase) CreateCourse(ctx context.Context, course *domain.Course) (int32, error) {
	user, err := uc.courseRepository.GetUser(ctx, course.AuthorID)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return 0, errors.NewNotFoundError("author not found", "author")
		}
		return 0, fmt.Errorf("getting user %d: %w", course.AuthorID, err)
	}

	if user.Type == domain.UserTypeTeacher {
		return 0, errors.NewInvalidInputError("Ученик не может создавать курсы.", "user_type")
	}

	id, err := uc.courseRepository.CreateCourse(ctx, course)
	if err != nil {
		return 0, fmt.Errorf("creating course: %w", err)
	}

	return id, nil
}

func (uc *CourseUseCase) GetCourse(ctx context.Context, id int32) (*domain.Course, error) {
	course, err := uc.courseRepository.GetCourse(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("Курс не найден.", "course")
		}
		return nil, fmt.Errorf("getting course %d: %w", id, err)
	}

	return course, nil
}

func (uc *CourseUseCase) GetAuthorCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	courses, err := uc.courseRepository.GetAuthorCourses(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("gettign author %d courses: %w", userID, err)
	}

	return courses, nil
}

func (uc *CourseUseCase) GetUserCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	courses, err := uc.courseRepository.GetUserCourses(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("gettign user %d courses: %w", userID, err)
	}

	return courses, nil
}

func (uc *CourseUseCase) UpdateCourse(ctx context.Context, course *domain.Course) error {
	course, err := uc.GetCourse(ctx, course.ID)
	if err != nil {
		return err
	}

	err = uc.courseRepository.UpdateCourse(ctx, course)
	if err != nil {
		return fmt.Errorf("updating course: %w", err)
	}

	return nil
}

func (uc *CourseUseCase) DeleteCourse(ctx context.Context, id int32) error {
	_, err := uc.GetCourse(ctx, id)
	if err != nil {
		return err
	}

	err = uc.courseRepository.DeleteCourse(ctx, id)
	if err != nil {
		return fmt.Errorf("deleting course: %w", err)
	}

	return nil
}

func (uc *CourseUseCase) GetCourseMembers(ctx context.Context, id int32) ([]domain.User, error) {
	_, err := uc.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	members, err := uc.courseRepository.GetCourseMembers(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting course members: %w", err)
	}

	return members, nil
}

func (uc *CourseUseCase) AddCourseMember(ctx context.Context, courseID int32, userID int64) error {
	_, err := uc.GetCourse(ctx, courseID)
	if err != nil {
		return err
	}

	_, err = uc.courseRepository.GetUser(ctx, userID)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return errors.NewNotFoundError("Пользователь не найден.", "user")
		}
		return fmt.Errorf("getting user %d: %w", userID, err)
	}

	err = uc.courseRepository.AddCourseMember(ctx, courseID, userID)
	if err != nil {
		return fmt.Errorf("adding course %d member %d: %w", courseID, userID, err)
	}

	return nil
}
