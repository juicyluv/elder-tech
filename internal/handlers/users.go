package handlers

import (
	stderrors "errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"
)

func FromDomainUserToUser(user *domain.User) *User {
	return &User{
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		Email:      user.Email,
		Gender:     user.Gender,
		Id:         user.ID,
		ImageId:    user.ImageID,
		LastOnline: user.LastOnline,
		Name:       user.Name,
		Patronymic: user.Patronymic,
		Phone:      user.Phone,
		Surname:    user.Surname,
	}
}

func (h HttpHandler) GetUser(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := db.GetUser(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Пользователь не найден.", "user"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting user %d: %w", id, err))
		return
	}

	render.JSON(w, r, FromDomainUserToUser(user))
}

func (h HttpHandler) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var req UpdateUserJSONRequestBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	user, err := db.GetUser(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Пользователь не найден.", "user"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting user %d: %w", user.ID, err))
		return
	}

	err = db.UpdateUser(r.Context(), &domain.User{
		ID:         id,
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		Age:        req.Age,
		Gender:     req.Gender,
		Phone:      req.Phone,
		Email:      req.Email,
	})
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("updating user %d: %w", user.ID, err))
		return
	}

	render.Status(r, http.StatusOK)
}

func (h HttpHandler) UpdateUserImage(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := db.GetUser(r.Context(), id)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, errors.NewNotFoundError("Пользователь не найден.", "user"))
			return
		}
		ErrorResponse(w, r, fmt.Errorf("getting user %d: %w", user.ID, err))
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}
	defer r.Body.Close()

	mt := mimetype.Detect(content)
	if !mt.Is("image/jpeg") &&
		!mt.Is("image/png") &&
		!mt.Is("image/webp") {
		ErrorResponse(w, r, errors.NewInvalidInputError("Некорректный mime type.", "mime_type"))
		return
	}

	filename := uuid.New().String() + mt.Extension()
	err = h.imageFileSys.Save(r.Context(), content, filename)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("saving image: %w", err))
		return
	}

	err = db.UpdateUserImage(r.Context(), user.ID, filename, user.ImageID)
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("updating user %d image: %w", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
