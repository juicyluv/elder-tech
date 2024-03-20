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

	Rating   float32
	Progress *float32

	CoverImage        *string
	CoverImageContent []byte

	Categories []CourseCategory
}

type CourseCategory struct {
	ID   int16
	Name string
}
