package http

import (
	stderrors "errors"
	"net/http"

	"github.com/go-chi/render"

	"diplom-backend/internal/common/errors"
)

type HttpHandler struct {
	userUseCase   UserUseCase
	authUseCase   AuthUseCase
	courseUseCase CourseUseCase
}

func NewHandler(
	userUseCase UserUseCase,
	authUseCase AuthUseCase,
	courseUseCase CourseUseCase,
) *HttpHandler {
	return &HttpHandler{
		userUseCase:   userUseCase,
		authUseCase:   authUseCase,
		courseUseCase: courseUseCase,
	}
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	var (
		domainError   errors.Error
		responseError Error
		statusCode    = http.StatusInternalServerError
	)

	if stderrors.As(err, &domainError) {
		responseError.Message = domainError.Error()
		responseError.Slug = domainError.Slug()

		switch domainError.Type() {
		case errors.ErrorTypeAuth:
			statusCode = http.StatusUnauthorized
		case errors.ErrorTypeNotFound:
			statusCode = http.StatusNotFound
		case errors.ErrorTypeInvalidInput:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}
	} else {
		responseError.Message = err.Error()
	}

	render.Status(r, statusCode)
	render.JSON(w, r, responseError)
}
