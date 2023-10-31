package errs

import "net/http"

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type ErrorData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: message,
		ErrError:   "NOT_AUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: message,
		ErrError:   "NOT_AUTHENTICATED",
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusInternalServerError, //500
		ErrMessage: message,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrMessage: message,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}
