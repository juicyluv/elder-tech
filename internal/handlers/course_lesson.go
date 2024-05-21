package handlers

import (
	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
)

func (h *HttpHandler) AddCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	var req AddCourseBlockLessonJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	courseBlock, err := db.GetCourseBlock(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Блок курса не найден.", "course_block"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block: %w", err))
		return
	}

	course, err := db.GetCourse(r.Context(), courseBlock.CourseID)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course: %w", err))
		return
	}

	ac, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting auth context: %w", err))
		return
	}

	if ac.ID != course.AuthorID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	lessonID, err := db.AddCourseBlockLesson(r.Context(), &domain.CourseBlockLesson{
		CourseBlockID: courseBlock.ID,
		Number:        req.Number,
		Title:         req.Title,
		Description:   req.Description,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, map[string]any{
		"id": lessonID,
	})
}

func (h *HttpHandler) GetCourseBlockLessons(w http.ResponseWriter, r *http.Request, id int64) {
	_, err := db.GetCourseBlock(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Блок курса не найден.", "course_block"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block: %w", err))
		return
	}

	lessons, err := db.GetCourseBlockLessons(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	courseBlockLessons := make([]CourseBlockLesson, 0, len(lessons))
	for _, lesson := range lessons {
		courseBlockLessons = append(courseBlockLessons, CourseBlockLesson{
			CourseBlockId: lesson.CourseBlockID,
			Description:   lesson.Description,
			Id:            lesson.ID,
			Number:        lesson.Number,
			Title:         lesson.Title,
		})
	}

	render.JSON(w, r, map[string]any{
		"lessons": courseBlockLessons,
	})
}

func (h *HttpHandler) UpdateCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	var req UpdateCourseBlockLessonJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	block, err := db.GetCourseBlock(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Блок курса не найден.", "course_block"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block %d: %w", id, err))
		return
	}

	course, err := db.GetCourse(r.Context(), block.CourseID)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Курс не найден.", "course"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
	}

	ac, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting auth context: %w", err))
		return
	}

	if ac.ID != course.AuthorID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = db.UpdateCourseBlockLesson(r.Context(), &domain.CourseBlockLesson{
		ID:          id,
		Number:      req.Number,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) DeleteCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	lesson, err := db.GetCourseBlockLesson(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "course_lesson"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course lesson %d: %w", id, err))
		return
	}

	courseBlock, err := db.GetCourseBlock(r.Context(), lesson.CourseBlockID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course block %d: %w", lesson.CourseBlockID, err))
		return
	}

	course, err := db.GetCourse(r.Context(), courseBlock.CourseID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", courseBlock.CourseID, err))
		return
	}

	ac, err := auth.GetAuthContextFromContext(r.Context())
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting auth context: %w", err))
		return
	}

	if ac.ID != course.AuthorID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = db.DeleteCourseBlockLesson(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
