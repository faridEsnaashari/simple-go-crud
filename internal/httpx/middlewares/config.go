package middlewares

import (
	"net/http"

	errorx "a.com/test/internal/httpx/errors"
)

func Config[T any](fn func(r *http.Request) (T, errorx.HandledError)) func(w http.ResponseWriter, r *http.Request) {
	withSafeCall := WithSafeCall(fn)
	withHandled := WithHandleError(withSafeCall)
	withRes := WithResponser(withHandled)
	withLog := WithLogger(withRes)

	return withLog
}
