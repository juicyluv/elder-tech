package http

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"diplom-backend/internal/domain"
)

type AuthUseCase interface {
	SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error)
	SignUp(ctx context.Context, req *domain.SignUpRequest) error
	GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error)
}

func (h HttpHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	response, err := h.authUseCase.SignIn(r.Context(), &domain.SignInRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SignInResponse{
		Id:         response.ID,
		Name:       response.Name,
		Patronymic: response.Patronymic,
		Surname:    response.Surname,
		Token:      response.Token,
	})
}

func (h HttpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err := h.authUseCase.SignUp(r.Context(), &domain.SignUpRequest{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Surname:  req.Surname,
		Email:    req.Email,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
