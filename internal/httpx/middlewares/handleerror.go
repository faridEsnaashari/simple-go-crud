package middlewares

import (
	"cmp"
	"net/http"

	errorx "a.com/test/internal/httpx/errors"
	logx "a.com/test/internal/httpx/logs"
)

func WithHandleError[T any](
	f func(r *http.Request) (T, errorx.HandledError),
) func(w http.ResponseWriter, r *http.Request) Response {
	return func(w http.ResponseWriter, r *http.Request) Response {
		result, err := f(r)

		if err.ErrorStatus == http.StatusInternalServerError {
			logx.Error("INTERNAL_ERROR", "err", err)
			return Response{
				Data:       "INTERNAL_ERROR",
				Message:    "INTERNAL_ERROR",
				Status:     FailedResponseStatus,
				HttpStatus: err.ErrorStatus,
			}
		}

		if err.ErrorStatus == http.StatusBadRequest {
			return Response{
				Data:       err.DtoErrors,
				Message:    cmp.Or(err.Message, "failed"),
				Status:     FailedResponseStatus,
				HttpStatus: err.ErrorStatus,
			}
		}

		if err.ErrorStatus != 0 {
			return Response{
				Data:       cmp.Or(err.Message, "failed"),
				Message:    cmp.Or(err.Message, "failed"),
				Status:     FailedResponseStatus,
				HttpStatus: err.ErrorStatus,
			}
		}

		return Response{
			Data:   result,
			Status: SuccessResponseStatus,
		}
	}
}
