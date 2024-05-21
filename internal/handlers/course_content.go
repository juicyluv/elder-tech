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

func (h *HttpHandler) AddCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	var req AddCourseBlockLessonContentJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	lesson, err := db.GetCourseBlockLesson(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "lesson"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block lesson: %w", err))
		return
	}

	courseBlock, err := db.GetCourseBlock(r.Context(), lesson.CourseBlockID)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "course_block"))
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

	contentID, err := db.AddCourseBlockLessonContent(r.Context(), &domain.CourseBlockLessonContent{
		LessonID: id,
		Number:   req.Number,
		Type:     req.Type,
		Value:    req.Value,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, map[string]any{
		"id": contentID,
	})
}

func (h *HttpHandler) GetCourseBlockLessonContents(w http.ResponseWriter, r *http.Request, id int64) {
	_, err := db.GetCourseBlockLesson(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "lesson"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block lesson: %w", err))
		return
	}

	contents, err := db.GetCourseBlockLessonContents(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	courseLessonContents := make([]CourseBlockLessonContent, 0, len(contents))
	for _, lesson := range contents {
		courseLessonContents = append(courseLessonContents, CourseBlockLessonContent{
			Id:       lesson.ID,
			LessonId: lesson.LessonID,
			Number:   lesson.Number,
			Type:     lesson.Type,
			Value:    lesson.Value,
		})
	}

	render.JSON(w, r, map[string]any{
		"lessons": courseLessonContents,
	})
}

func (h *HttpHandler) UpdateCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	var req UpdateCourseBlockLessonContentJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	content, err := db.GetCourseBlockLessonContent(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Контент урока не найден.", "lesson_content"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block lesson content %d: %w", id, err))
		return
	}

	lesson, err := db.GetCourseBlockLesson(r.Context(), content.LessonID)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "lesson"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block lesson %d: %w", content.LessonID, err))
		return
	}

	block, err := db.GetCourseBlock(r.Context(), lesson.CourseBlockID)
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

	err = db.UpdateCourseBlockLessonContent(r.Context(), &domain.CourseBlockLessonContent{
		ID:     id,
		Number: req.Number,
		Value:  req.Value,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) DeleteCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	content, err := db.GetCourseBlockLessonContent(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Контент урока не найден.", "lesson_content"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course block lesson content %d: %w", id, err))
		return
	}

	lesson, err := db.GetCourseBlockLesson(r.Context(), content.LessonID)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Урок не найден.", "course_lesson"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting course lesson %d: %w", content.LessonID, err))
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

	err = db.DeleteCourseBlockLessonContent(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
