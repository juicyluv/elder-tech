package http

import (
	"encoding/json"
	stderrors "errors"
	"net/http"

	"diplom-backend/internal/common/errors"
)

type HttpHandler struct {
	userUseCase UserUseCase
}

func NewHandler(userUseCase UserUseCase) *HttpHandler {
	return &HttpHandler{userUseCase: userUseCase}
}

func ErrorResponse(w http.ResponseWriter, err error) {
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

	response, err := json.Marshal(&responseError)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response, _ = json.Marshal(Error{
			Message: "unknown error",
		})
		_, _ = w.Write(response)
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
