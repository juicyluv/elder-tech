package handlers

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
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

func (h HttpHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req CreateCourseJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// id, err := db.CreateCourse(r.Context(), course)
	// if err != nil {
	// 	return 0, fmt.Errorf("creating course: %w", err)
	// }

	// return id, nil
}

func (h HttpHandler) GetAuthorCourses(w http.ResponseWriter, r *http.Request, id int64) {
	courses, err := db.GetAuthorCourses(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("gettign author %d courses: %w", id, err))
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
	courses, err := db.GetUserCourses(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("gettign user %d courses: %w", id, err))
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
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if err != nil {
			if stderrors.Is(err, pgx.ErrNoRows) {
				ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
				return
			}
			ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
			return
		}
	}

	err = db.DeleteCourse(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("deleting course: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h HttpHandler) GetCourse(w http.ResponseWriter, r *http.Request, id int32) {
	course, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if err != nil {
			if stderrors.Is(err, pgx.ErrNoRows) {
				ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
				return
			}
			ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
			return
		}
	}

	render.JSON(w, r, FromDomainCourseToCourse(course))
}

func (h HttpHandler) UpdateCourse(w http.ResponseWriter, r *http.Request, id int32) {
	// TODO implement me
	panic("implement me")

	// course, err := uc.GetCourse(r.Context(), course.ID)
	// if err != nil {
	// 	return err
	// }

	// err = db.UpdateCourse(r.Context(), course)
	// if err != nil {
	// 	return fmt.Errorf("updating course: %w", err)
	// }
}

func (h HttpHandler) GetCourseMembers(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if err != nil {
			if stderrors.Is(err, pgx.ErrNoRows) {
				ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
				return
			}
			ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
			return
		}
	}

	members, err := db.GetCourseMembers(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course members: %w", err))
		return
	}

	users := make([]*User, 0, len(members))
	for _, user := range members {
		users = append(users, FromDomainUserToUser(&user))
	}

	render.JSON(w, r, map[string]any{
		"members": users,
	})
}

func (h HttpHandler) AddCourseMembers(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if err != nil {
			if stderrors.Is(err, pgx.ErrNoRows) {
				ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
				return
			}
			ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
			return
		}
	}

	user, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting auth context"))
		return
	}

	err = db.AddCourseMember(r.Context(), id, user.ID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("adding course %d member %d: %w", id, user.ID, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
