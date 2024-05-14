package handlers

import (
	stderrors "errors"
	"fmt"
	"net/http"
	"time"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
)

func FromDomainCourseToCourse(course *domain.Course) *Course {
	c := &Course{
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
		Categories:          course.Categories,
	}

	if len(course.Blocks) != 0 {
		blocks := make([]CourseBlock, 0, len(course.Blocks))
		for _, block := range course.Blocks {
			blocks = append(blocks, CourseBlock{
				CourseId:    block.CourseID,
				Description: block.Description,
				Id:          block.ID,
				Number:      block.Number,
				Title:       block.Description,
			})
		}
		c.CourseBlocks = &blocks
	}

	return c
}

func (h HttpHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req CreateCourseJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	user, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	id, err := db.CreateCourse(r.Context(), &domain.Course{
		AuthorID:              user.ID,
		Title:                 req.Title,
		Description:           req.Description,
		Difficulty:            req.Difficulty,
		TimeToCompleteMinutes: req.TimeToCompleteMinutes,
		About:                 req.About,
		ForWho:                req.ForWho,
		Requirements:          req.Requirements,
		CreatedAt:             time.Now(),
		Categories:            req.Categories,
	})
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("creating course: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]any{
		"id": id,
	})
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
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
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
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
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
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
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

func (h HttpHandler) JoinCourse(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
	}

	user, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err = db.AddCourseMember(r.Context(), id, user.ID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("adding course %d member %d: %w", id, user.ID, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h HttpHandler) LeaveCourse(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
	}

	user, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err = db.RemoveCourseMember(r.Context(), id, user.ID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("removing course %d member %d: %w", id, user.ID, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h HttpHandler) GetCourseCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := db.GetCourseCategories(r.Context())
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course categores: %w", err))
		return
	}

	responseCategories := make([]CourseCategory, 0, len(categories))
	for _, c := range categories {
		responseCategories = append(responseCategories, CourseCategory{
			Id:   c.ID,
			Name: c.Name,
		})
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]any{
		"categories": responseCategories,
	})
}
