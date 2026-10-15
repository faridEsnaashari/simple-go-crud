package middlewares

import (
	"net/http"

	errorx "a.com/test/internal/httpx/errors"
)

func Config[T any](fn func(r Request) (T, errorx.HandledError)) func(w http.ResponseWriter, r *http.Request) {
	withReq := WithRequest(fn)
	withValid := WithHandleError(withReq)
	withRes := WithResponser(withValid)

	return withRes
}
