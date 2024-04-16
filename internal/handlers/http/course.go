package http

import (
	"context"
	"net/http"

	"diplom-backend/internal/domain"

	"github.com/go-chi/render"
)

func FromDomainCourseToCourse(course *domain.Course) *Course {
	return &Course{
		About:               course.About,
		AuthorId:            course.AuthorID,
		CoverImage:          course.CoverImage,
		CreatedAt:           course.CreatedAt,
		Description:         course.Description,
		Difficulty:          course.Difficulty,
		ForWho:              course.ForWho,
		Id:                  course.ID,
		Progress:            course.Progress,
		Rating:              course.Rating,
		Requirements:        course.Requirements,
		TimeToCompleteHours: course.TimeToCompleteMinutes,
		Title:               course.Title,
		UpdatedAt:           course.UpdatedAt,
	}
}

type CourseUseCase interface {
	GetCourse(ctx context.Context, id int32) (*domain.Course, error)
	GetAuthorCourses(ctx context.Context, userID int64) ([]domain.Course, error)
	GetUserCourses(ctx context.Context, userID int64) ([]domain.Course, error)
}

func (h HttpHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h HttpHandler) GetAuthorCourses(w http.ResponseWriter, r *http.Request, id int64) {
	courses, err := h.courseUseCase.GetAuthorCourses(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	httpCourses := make([]*Course, 0, len(courses))
	for _, course := range courses {
		httpCourses = append(httpCourses, FromDomainCourseToCourse(&course))
	}

	render.JSON(w, r, map[string]any{
		"courses": httpCourses,
	})
}

func (h HttpHandler) GetUserCourses(w http.ResponseWriter, r *http.Request, id int64) {
	courses, err := h.courseUseCase.GetUserCourses(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	httpCourses := make([]*Course, 0, len(courses))
	for _, course := range courses {
		httpCourses = append(httpCourses, FromDomainCourseToCourse(&course))
	}

	render.JSON(w, r, map[string]any{
		"courses": httpCourses,
	})
}

func (h HttpHandler) DeleteCourse(w http.ResponseWriter, r *http.Request, id int32) {
	// TODO implement me
	panic("implement me")
}

func (h HttpHandler) GetCourse(w http.ResponseWriter, r *http.Request, id int32) {
	course, err := h.courseUseCase.GetCourse(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, FromDomainCourseToCourse(course))
}

func (h HttpHandler) UpdateCourse(w http.ResponseWriter, r *http.Request, id int32) {
	// TODO implement me
	panic("implement me")
}

func (h HttpHandler) GetCourseMembers(w http.ResponseWriter, r *http.Request, id int64) {
	// TODO implement me
	panic("implement me")
}

func (h HttpHandler) AddCourseMembers(w http.ResponseWriter, r *http.Request, id int64) {
	// TODO implement me
	panic("implement me")
}
