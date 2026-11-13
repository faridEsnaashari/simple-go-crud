package middlewares

import (
	"net/http"
	"runtime/debug"

	errorx "a.com/test/internal/httpx/errors"
	logx "a.com/test/internal/httpx/logs"
)

func WithSafeCall[T any](
	f func(r *http.Request) (res T, herr errorx.HandledError),
) func(r *http.Request) (res T, herr errorx.HandledError) {
	return func(r *http.Request) (res T, herr errorx.HandledError) {
		defer func() {
			if rec := recover(); rec != nil {
				logx.Error("PANIC", "panic", rec, "stack", string(debug.Stack()))
				herr = errorx.InternalError(nil)
			}
		}()

		return f(r)
	}
}
