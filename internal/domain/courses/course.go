package courses

type Course struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    int64  `json:"authorID"`
}
