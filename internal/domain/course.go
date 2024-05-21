package domain

import "time"

type Course struct {
	ID       int32
	AuthorID int64

	Title                 string
	Description           string
	Difficulty            int16
	TimeToCompleteMinutes int16
	About                 string
	ForWho                string
	Requirements          string

	CreatedAt time.Time
	UpdatedAt *time.Time

	Rating     *float32
	Progress   *float32
	CoverImage *int64

	Categories []int16
	Blocks     []CourseBlock

	AuthorName       string
	AuthorSurname    *string
	AuthorPatronymic *string
}

func (c *Course) CalculateRating(ratingSum, ratingCount *int) {
	if ratingCount == nil && ratingSum == nil {
		return
	}

	rating := float32(*ratingSum) / float32(*ratingCount)
	c.Rating = &rating
}

type CourseCategory struct {
	ID   int16
	Name string
}

type CourseBlock struct {
	ID          int64
	CourseID    int32
	Number      int
	Title       string
	Description string
}
