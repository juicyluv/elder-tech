package db

import (
	"context"
	"diplom-backend/internal/domain"
	"fmt"
)

func AddCourseBlock(ctx context.Context, block *domain.CourseBlock) (int64, error) {
	var id int64
	err := db.QueryRow(ctx, `
		INSERT INTO course_blocks(course_id,number,title,description)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		block.CourseID,
		block.Number,
		block.Title,
		block.Description,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting course block: %w", err)
	}

	return id, nil
}

func GetCourseBlock(ctx context.Context, id int64) (*domain.CourseBlock, error) {
	var block domain.CourseBlock
	err := db.QueryRow(ctx, `
		SELECT id,course_id,number,title,description
		FROM course_blocks
		WHERE id=$1
	`,
		id,
	).Scan(
		&block.ID,
		&block.CourseID,
		&block.Number,
		&block.Title,
		&block.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting course block: %w", err)
	}

	return &block, nil
}

func GetCourseBlocks(ctx context.Context, courseID int32) ([]domain.CourseBlock, error) {
	rows, err := db.Query(ctx, `
		SELECT id,course_id,number,title,description
		FROM course_blocks
		WHERE course_id=$1
		ORDER BY number
	`,
		courseID,
	)
	if err != nil {
		return nil, fmt.Errorf("selecting course blocks: %w", err)
	}
	defer rows.Close()

	var blocks []domain.CourseBlock
	var block domain.CourseBlock
	for rows.Next() {
		if err = rows.Scan(
			&block.ID,
			&block.CourseID,
			&block.Number,
			&block.Title,
			&block.Description,
		); err != nil {
			return nil, fmt.Errorf("scanning course block: %w", err)
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func UpdateCourseBlock(ctx context.Context, block *domain.CourseBlock) error {
	_, err := db.Exec(ctx, `
		UPDATE course_blocks
		SET number=$2,title=$3,description=$4
		WHERE id=$1
	`,
		block.ID,
		block.Number,
		block.Title,
		block.Description,
	)
	if err != nil {
		return fmt.Errorf("updating course block: %w", err)
	}

	return nil
}

func DeleteCourseBlock(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, `
		DELETE FROM course_blocks
		WHERE id=$1
	`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting course block: %w", err)
	}

	return nil
}
