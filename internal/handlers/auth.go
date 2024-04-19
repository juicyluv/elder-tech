package handlers

import (
	stderrors "errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"

	"diplom-backend/internal/common/auth"
	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/db"
	"diplom-backend/internal/domain"
)

func (h HttpHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	user, err := db.GetUserByPhone(r.Context(), req.Phone)
	if err != nil && !stderrors.Is(err, pgx.ErrNoRows) {
		ErrorResponse(w, r, err)
		return
	}

	if stderrors.Is(err, pgx.ErrNoRows) || !user.ComparePassword(req.Password) {
		ErrorResponse(w, r, errors.NewAuthError("Неправильный телефон или пароль.", "credentials"))
		return
	}

	token, err := auth.GenerateJWT(auth.Claims{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		ErrorResponse(w, r, fmt.Errorf("generating token: %w", err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SignInResponse{
		Id:         user.ID,
		Name:       user.Name,
		Patronymic: user.Patronymic,
		Surname:    user.Surname,
		Token:      token,
	})
}

func (h HttpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err := db.CheckPhoneUnique(r.Context(), req.Phone)
	if err != nil && !stderrors.Is(err, pgx.ErrNoRows) {
		ErrorResponse(w, r, fmt.Errorf("checking phone unique: %w", err))
		return
	}
	if err == nil {
		ErrorResponse(w, r, errors.NewInvalidInputError("Телефон уже используется.", "phone"))
		return
	}

	if req.Email != nil {
		err = db.CheckEmailUnique(r.Context(), *req.Email)
		if err != nil && !stderrors.Is(err, pgx.ErrNoRows) {
			ErrorResponse(w, r, fmt.Errorf("checking email unique: %w", err))
			return
		}
		if err == nil {
			ErrorResponse(w, r, errors.NewInvalidInputError("Почта уже используется.", "email"))
			return
		}
	}

	user := domain.User{
		Name:              req.Name,
		Phone:             req.Phone,
		PasswordEncrypted: req.Password,
		CreatedAt:         time.Now(),
		Surname:           req.Surname,
		Email:             req.Email,
	}
	if err = user.EncryptPassword(); err != nil {
		ErrorResponse(w, r, fmt.Errorf("encrypting password: %w", err))
		return
	}

	if err = db.CreateUser(r.Context(), &user); err != nil {
		ErrorResponse(w, r, fmt.Errorf("creating user: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
