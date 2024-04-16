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

	Rating   *float32
	Progress *float32

	CoverImage        *string
	CoverImageContent []byte

	Categories []CourseCategory

	AuthorName       string
	AuthorSurname    string
	AuthorPatronymic string
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
