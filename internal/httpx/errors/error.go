package errorx

import (
	"fmt"
	"net/http"

	dtox "a.com/test/internal/httpx/dto"
)

type HandledError struct {
	ErrorStatus int                        `json:"errorStatus"`
	Reason      string                     `json:"reason"`
	Message     string                     `json:"message"`
	DtoErrors   []dtox.DtoValidationFailed `json:"dtoErrors,omitempty"`
}

func NotFoundError(ctx string) HandledError {
	return HandledError{
		ErrorStatus: http.StatusNotFound,
		Reason:      fmt.Sprintf("%v not found", ctx),
		Message:     fmt.Sprintf("%v not found", ctx),
	}
}

func ValidationError(errors []dtox.DtoValidationFailed) HandledError {
	return HandledError{
		ErrorStatus: http.StatusBadRequest,
		Reason:      "validation error",
		Message:     "validation error",
		DtoErrors:   errors,
	}
}

func DatabaseError(err error) HandledError {
	return HandledError{
		ErrorStatus: http.StatusInternalServerError,
		Reason:      "database error",
		Message:     err.Error(),
	}
}

func InternalError(err error) HandledError {
	message := "internal"
	if err != nil {
		message = err.Error()
	}
	return HandledError{
		ErrorStatus: http.StatusInternalServerError,
		Reason:      "internal",
		Message:     message,
	}
}
