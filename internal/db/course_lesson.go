package db

import (
	"context"
	"diplom-backend/internal/domain"
	"fmt"
)

func AddCourseBlockLesson(ctx context.Context, block *domain.CourseBlockLesson) (int64, error) {
	var id int64
	err := db.QueryRow(ctx, `
		INSERT INTO course_block_lessons(course_block_id,number,title,description)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		block.CourseBlockID,
		block.Number,
		block.Title,
		block.Description,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting course block lesson: %w", err)
	}

	return id, nil
}

func GetCourseBlockLesson(ctx context.Context, id int64) (*domain.CourseBlockLesson, error) {
	var lesson domain.CourseBlockLesson
	err := db.QueryRow(ctx, `
		SELECT id,course_block_id,number,title,description
		FROM course_block_lessons
		WHERE id=$1
	`,
		id,
	).Scan(
		&lesson.ID,
		&lesson.CourseBlockID,
		&lesson.Number,
		&lesson.Title,
		&lesson.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting course block lesson: %w", err)
	}

	return &lesson, nil
}

func GetCourseBlockLessons(ctx context.Context, blockID int64) ([]domain.CourseBlockLesson, error) {
	rows, err := db.Query(ctx, `
		SELECT id,course_block_id,number,title,description
		FROM course_block_lessons
		WHERE course_block_id=$1
		ORDER BY number
	`,
		blockID,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting course block lessons: %w", err)
	}
	defer rows.Close()

	var lessons []domain.CourseBlockLesson
	var lesson domain.CourseBlockLesson
	for rows.Next() {
		if err = rows.Scan(
			&lesson.ID,
			&lesson.CourseBlockID,
			&lesson.Number,
			&lesson.Title,
			&lesson.Description,
		); err != nil {
			return nil, fmt.Errorf("scanning course block lesson: %w", err)
		}

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

func UpdateCourseBlockLesson(ctx context.Context, lesson *domain.CourseBlockLesson) error {
	_, err := db.Exec(ctx, `
		UPDATE course_block_lessons
		SET number=$2,title=$3,description=$4
		WHERE id=$1
	`,
		lesson.ID,
		lesson.Number,
		lesson.Title,
		lesson.Description,
	)
	if err != nil {
		return fmt.Errorf("updating course block lesson: %w", err)
	}

	return nil
}

func DeleteCourseBlockLesson(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, `
		DELETE FROM course_block_lessons
		WHERE id=$1
	`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting course block lesson: %w", err)
	}

	return nil
}
