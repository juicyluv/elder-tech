package db

import (
	"context"
	"diplom-backend/internal/domain"
	"fmt"
)

func AddCourseBlockLessonContent(ctx context.Context, block *domain.CourseBlockLessonContent) (int64, error) {
	var id int64
	err := db.QueryRow(ctx, `
		INSERT INTO course_block_lesson_content(course_block_lesson_id,number,type,value)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		block.LessonID,
		block.Number,
		block.Type,
		block.Value,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting content: %w", err)
	}

	return id, nil
}

func GetCourseBlockLessonContent(ctx context.Context, id int64) (*domain.CourseBlockLessonContent, error) {
	var lesson domain.CourseBlockLessonContent
	err := db.QueryRow(ctx, `
		SELECT id,course_block_lesson_id,number,type,value
		FROM course_block_lesson_content
		WHERE id=$1
	`,
		id,
	).Scan(
		&lesson.ID,
		&lesson.LessonID,
		&lesson.Number,
		&lesson.Type,
		&lesson.Value,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting content: %w", err)
	}

	return &lesson, nil
}

func GetCourseBlockLessonContents(ctx context.Context, blockID int64) ([]domain.CourseBlockLessonContent, error) {
	rows, err := db.Query(ctx, `
		SELECT id,course_block_lesson_id,number,type,value
		FROM course_block_lesson_content
		WHERE course_block_lesson_id=$1
		ORDER BY number
	`,
		blockID,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting content: %w", err)
	}
	defer rows.Close()

	var lessons []domain.CourseBlockLessonContent
	var lesson domain.CourseBlockLessonContent
	for rows.Next() {
		if err = rows.Scan(
			&lesson.ID,
			&lesson.LessonID,
			&lesson.Number,
			&lesson.Type,
			&lesson.Value,
		); err != nil {
			return nil, fmt.Errorf("scanning content: %w", err)
		}

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

func UpdateCourseBlockLessonContent(ctx context.Context, lesson *domain.CourseBlockLessonContent) error {
	_, err := db.Exec(ctx, `
		UPDATE course_block_lesson_content
		SET number=$2,value=$3
		WHERE id=$1
	`,
		lesson.ID,
		lesson.Number,
		lesson.Value,
	)
	if err != nil {
		return fmt.Errorf("updating content: %w", err)
	}

	return nil
}

func DeleteCourseBlockLessonContent(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, `
		DELETE FROM course_block_lesson_content
		WHERE id=$1
	`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting content: %w", err)
	}

	return nil
}
