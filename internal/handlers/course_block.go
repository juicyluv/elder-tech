package handlers

import (
	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (h *HttpHandler) AddCourseBlock(w http.ResponseWriter, r *http.Request, id int32) {
	var req AddCourseBlockJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	course, err := db.GetCourse(r.Context(), id)
	if err != nil {
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

	blockID, err := db.AddCourseBlock(r.Context(), &domain.CourseBlock{
		CourseID:    id,
		Number:      req.Number,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, map[string]any{
		"id": blockID,
	})
}

func (h *HttpHandler) GetCourseBlocks(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := db.GetCourse(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	blocks, err := db.GetCourseBlocks(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	courseBlocks := make([]CourseBlock, 0, len(blocks))
	for _, block := range blocks {
		courseBlocks = append(courseBlocks, CourseBlock{
			CourseId:    block.CourseID,
			Description: block.Description,
			Id:          block.ID,
			Number:      block.Number,
			Title:       block.Title,
		})
	}

	render.JSON(w, r, map[string]any{
		"blocks": courseBlocks,
	})
}

func (h *HttpHandler) UpdateCourseBlock(w http.ResponseWriter, r *http.Request, id int64) {
	var req AddCourseBlockJSONBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	block, err := db.GetCourseBlock(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
	}

	course, err := db.GetCourse(r.Context(), block.CourseID)
	if err != nil {
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

	err = db.UpdateCourseBlock(r.Context(), &domain.CourseBlock{
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

func (h *HttpHandler) DeleteCourseBlock(w http.ResponseWriter, r *http.Request, id int64) {
	block, err := db.GetCourseBlock(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("getting course %d: %w", id, err))
		return
	}

	course, err := db.GetCourse(r.Context(), block.CourseID)
	if err != nil {
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

	err = db.DeleteCourseBlock(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
