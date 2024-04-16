package postgresql

import (
	"context"
	"errors"
	"fmt"

	"diplom-backend/internal/domain"
	"diplom-backend/internal/infrastructure/repository"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetCourse(ctx context.Context, id int32) (*domain.Course, error) {
	var course domain.Course
	var ratingCount, ratingSum *int

	err := r.db.QueryRow(ctx, `
SELECT c.id,
       c.title,
       c.description,
       c.difficulty,
       c.time_to_complete_minutes,
       c.about,
       c.for_who,
       c.requirements,
       c.created_at,
       c.updated_at,
       c.cover_image,
       c.author_id,
       u.name,
       u.surname,
       u.patronymic,
       ratings.num,
       ratings.rating
FROM courses c
JOIN users u ON u.id=c.id
LEFT JOIN
    (SELECT course_id,
            count(*) AS num,
            sum(rating) AS rating
     FROM course_ratings
     WHERE course_id = $1
     GROUP BY course_id) ratings ON ratings.course_id=c.id
WHERE c.id=$1`, id).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.Difficulty,
		&course.TimeToCompleteMinutes,
		&course.About,
		&course.ForWho,
		&course.Requirements,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.CoverImage,

		&course.AuthorID,
		&course.AuthorName,
		&course.AuthorSurname,
		&course.AuthorPatronymic,

		&ratingCount,
		&ratingSum,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrNotFound
		}

		return nil, fmt.Errorf("selecting course: %w", err)
	}

	course.CalculateRating(ratingSum, ratingCount)

	return &course, nil
}

func (r *Repository) GetAuthorCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	rows, err := r.db.Query(ctx, `
SELECT c.id,
       c.title,
       c.description,
       c.difficulty,
       c.time_to_complete_minutes,
       c.about,
       c.for_who,
       c.requirements,
       c.created_at,
       c.updated_at,
       c.cover_image,
       c.author_id,
       ratings.num,
       ratings.rating
FROM courses c
LEFT JOIN
    (SELECT course_id,
            count(*) AS num,
            sum(rating) AS rating
     FROM course_ratings
     GROUP BY course_id) ratings ON ratings.course_id = c.id
WHERE c.author_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("selecting courses: %w", err)
	}
	defer rows.Close()

	courses := make([]domain.Course, 0, 16)
	for rows.Next() {
		var course domain.Course
		var rating, ratingCount *int

		err = rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Difficulty,
			&course.TimeToCompleteMinutes,
			&course.About,
			&course.ForWho,
			&course.Requirements,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.CoverImage,
			&course.AuthorID,

			&rating,
			&ratingCount,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning course: %w", err)
		}

		course.CalculateRating(rating, ratingCount)

		courses = append(courses, course)
	}

	return courses, nil
}

func (r *Repository) GetUserCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	rows, err := r.db.Query(ctx, `
SELECT c.id,
       c.title,
       c.description,
       c.difficulty,
       c.time_to_complete_minutes,
       c.about,
       c.for_who,
       c.requirements,
       c.created_at,
       c.updated_at,
       c.cover_image,
       ratings.num,
       ratings.rating,
       c.author_id,
       u.name,
       u.surname,
       u.patronymic
FROM courses c
JOIN
    (SELECT course_id
     FROM course_members
     WHERE user_id = $1 ) mem ON mem.course_id = c.id
JOIN users u ON u.id=c.author_id
LEFT JOIN
    (SELECT course_id,
            count(*) AS num,
            sum(rating) AS rating
     FROM course_ratings
     GROUP BY course_id) ratings ON ratings.course_id = c.id`, userID)
	if err != nil {
		return nil, fmt.Errorf("selecting courses: %w", err)
	}
	defer rows.Close()

	courses := make([]domain.Course, 0, 16)
	for rows.Next() {
		var course domain.Course
		var rating, ratingCount *int

		err = rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Difficulty,
			&course.TimeToCompleteMinutes,
			&course.About,
			&course.ForWho,
			&course.Requirements,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.CoverImage,

			&rating,
			&ratingCount,

			&course.AuthorID,
			&course.AuthorName,
			&course.AuthorSurname,
			&course.AuthorPatronymic,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning course: %w", err)
		}

		course.CalculateRating(rating, ratingCount)

		courses = append(courses, course)
	}

	return courses, nil
}

func (r *Repository) CreateCourse(ctx context.Context, course *domain.Course) (int32, error) {
	return 0, nil
}

func (r *Repository) UpdateCourse(ctx context.Context, course *domain.Course) error {
	return nil
}

func (r *Repository) DeleteCourse(ctx context.Context, id int32) error {
	return nil
}

func (r *Repository) GetCourseMembers(ctx context.Context, id int32) ([]domain.User, error) {
	return nil, nil
}

func (r *Repository) AddCourseMember(ctx context.Context, courseID int32, userID int64) error {
	return nil
}
