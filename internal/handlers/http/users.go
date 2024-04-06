package http

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"diplom-backend/internal/domain"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

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
	user, err := h.userUseCase.GetUser(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
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

	err = h.userUseCase.UpdateUser(r.Context(), &domain.User{
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
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
}
