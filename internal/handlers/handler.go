package handlers

import (
	stderrors "errors"
	"net/http"

	"github.com/go-chi/render"

	"diplom-backend/internal/common/errors"
	"diplom-backend/internal/filesystem"
)

type HttpHandler struct {
	imageFileSys *filesystem.FileSystem
}

func NewHandler(
	imageFileSys *filesystem.FileSystem,
) *HttpHandler {
	return &HttpHandler{
		imageFileSys: imageFileSys,
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
