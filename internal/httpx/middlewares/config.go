package middlewares

import (
	"net/http"

	errorx "a.com/test/internal/httpx/errors"
)

func Config[T any](fn func(r *http.Request) (T, errorx.HandledError)) func(w http.ResponseWriter, r *http.Request) {
	withValid := WithHandleError(fn)
	withRes := WithResponser(withValid)
	withLog := WithLogger(withRes)

	return withLog
}
