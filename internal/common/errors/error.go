package errors

type ErrorType string

const (
	ErrorTypeInvalidInput ErrorType = "invalid_input"
	ErrorTypeNotFound     ErrorType = "not_found"
	ErrorTypeAuth         ErrorType = "auth"
	ErrorTypeInternal     ErrorType = "internal"
)

type Error struct {
	error     string
	slug      string
	errorType ErrorType
}

func (e Error) Error() string {
	return e.error
}

func (e Error) Slug() string {
	return e.slug
}

func (e Error) Type() ErrorType {
	return e.errorType
}

func NewInvalidInputError(error, slug string) Error {
	return Error{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeInvalidInput,
	}
}

func NewNotFoundError(error, slug string) Error {
	return Error{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeNotFound,
	}
}

func NewAuthError(error, slug string) Error {
	return Error{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeAuth,
	}
}

func NewInternalError(error, slug string) Error {
	return Error{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeInternal,
	}
}
