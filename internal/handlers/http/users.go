package http

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"diplom-backend/internal/domain/users"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*users.User, error)
}

func (h HttpHandler) GetUser(w http.ResponseWriter, r *http.Request, id string) {
	user, err := h.userUseCase.GetUser(r.Context(), id)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	render.JSON(w, r, user)
}
