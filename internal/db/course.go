package db

import (
	"context"
	"fmt"
	"time"

	"diplom-backend/internal/domain"
)

func GetCourse(ctx context.Context, id int32) (*domain.Course, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("beginning tx")
	}
	defer tx.Rollback(context.Background())

	var course domain.Course
	var ratingCount, ratingSum *int

	err = tx.QueryRow(ctx, `
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
		WHERE c.id=$1`,
		id,
	).Scan(
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
		return nil, fmt.Errorf("selecting course: %w", err)
	}

	course.CalculateRating(ratingSum, ratingCount)

	course.Categories, err = getCourseCategories(ctx, course.ID)
	if err != nil {
		return nil, fmt.Errorf("getting course categories: %w", err)
	}

	course.Blocks, err = getCourseBlocks(ctx, course.ID)
	if err != nil {
		return nil, fmt.Errorf("getting course blocks: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("committing tx: %w", err)
	}

	return &course, nil
}

func GetAuthorCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("beginning tx")
	}
	defer tx.Rollback(context.Background())

	rows, err := tx.Query(ctx, `
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

		course.Categories, err = getCourseCategories(ctx, course.ID)
		if err != nil {
			return nil, fmt.Errorf("getting course categories: %w", err)
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func GetUserCourses(ctx context.Context, userID int64) ([]domain.Course, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("beginning tx")
	}
	defer tx.Rollback(context.Background())

	rows, err := tx.Query(ctx, `
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
			 WHERE user_id = $1) mem ON mem.course_id = c.id
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

		course.Categories, err = getCourseCategories(ctx, course.ID)
		if err != nil {
			return nil, fmt.Errorf("getting course categories: %w", err)
		}

		courses = append(courses, course)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("committing tx: %w", err)
	}

	return courses, nil
}

func CreateCourse(ctx context.Context, course *domain.Course) (int32, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("beginning tx: %w", err)
	}
	defer tx.Rollback(context.Background())
	var id int32

	err = tx.QueryRow(ctx, `
		INSERT INTO courses(title, description, difficulty, time_to_complete_minutes, about, for_who, requirements, created_at, updated_at, cover_image, author_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`,
		course.Title,
		course.Description,
		course.Difficulty,
		course.TimeToCompleteMinutes,
		course.About,
		course.ForWho,
		course.Requirements,
		time.Now(),
		time.Now(),
		course.CoverImage,
		course.AuthorID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting course: %w", err)
	}

	for _, c := range course.Categories {
		_, err = tx.Exec(ctx, "INSERT INTO courses_to_categories(course_id, course_category_id) VALUES ($1, $2)", id, c)
		if err != nil {
			return 0, fmt.Errorf("inserting categories: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("committing tx")
	}

	return id, nil
}

func UpdateCourse(ctx context.Context, course *domain.Course) error {
	return nil
}

func DeleteCourse(ctx context.Context, id int32) error {
	_, err := db.Exec(ctx, `DELETE FROM courses WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("deleting course: %w", err)
	}

	return nil
}

func GetCourseMembers(ctx context.Context, id int32) ([]domain.User, error) {
	rows, err := db.Query(ctx, `
		SELECT
			u.id,
			u.name,
			u.surname,
			u.patronymic,
			u.image_id,
			u.last_online
		FROM course_members cm
		JOIN users u ON u.id = cm.user_id
		WHERE cm.course_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("selecting course members: %w", err)
	}
	defer rows.Close()

	members := make([]domain.User, 0, 20)
	for rows.Next() {
		var user domain.User

		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.ImageID,
			&user.LastOnline,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning course member: %w", err)
		}

		members = append(members, user)
	}

	return members, nil
}

func AddCourseMember(ctx context.Context, courseID int32, userID int64) error {
	_, err := db.Exec(ctx, `
		INSERT INTO course_members(course_id, user_id)
		VALUES ($1, $2)`,
		courseID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("inserting course member: %w", err)
	}

	return nil
}

func RemoveCourseMember(ctx context.Context, courseID int32, userID int64) error {
	_, err := db.Exec(ctx, `
		DELETE FROM course_members
		WHERE course_id = $1 AND user_id = $2`,
		courseID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("deleting course member: %w", err)
	}

	return nil
}

func GetCourseCategories(ctx context.Context) ([]domain.CourseCategory, error) {
	rows, err := db.Query(ctx, "SELECT id, name FROM course_categories ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("selecting categories: %w", err)
	}
	defer rows.Close()

	categories := make([]domain.CourseCategory, 0, 32)
	var c domain.CourseCategory
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, fmt.Errorf("scanning category: %w", err)
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func getCourseCategories(ctx context.Context, id int32) ([]int16, error) {
	rows, err := db.Query(ctx, "SELECT course_category_id FROM courses_to_categories WHERE course_id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("selecting course categories: %w", err)
	}
	defer rows.Close()

	categories := make([]int16, 0, 8)
	var category int16
	for rows.Next() {
		if err = rows.Scan(&category); err != nil {
			return nil, fmt.Errorf("scanning category: %w", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func getCourseBlocks(ctx context.Context, id int32) ([]domain.CourseBlock, error) {
	rows, err := db.Query(ctx, `
		SELECT id, course_id, number, title, description
		FROM course_blocks
		WHERE course_id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("selecting course blocks: %w", err)
	}
	defer rows.Close()

	blocks := make([]domain.CourseBlock, 0, 8)
	var block domain.CourseBlock
	for rows.Next() {
		if err = rows.Scan(
			&block.ID,
			&block.CourseID,
			&block.Number,
			&block.Title,
			&block.Description,
		); err != nil {
			return nil, fmt.Errorf("scanning block: %w", err)
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
