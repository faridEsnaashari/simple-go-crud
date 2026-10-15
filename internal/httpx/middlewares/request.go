package middlewares

import (
	"fmt"
	"net/http"

	errorx "a.com/test/internal/httpx/errors"
)

type Request struct {
	*http.Request
}

func (r *Request) Logger() {
	fmt.Println(r.Method + " -> " + r.URL.String())
}

func WithRequest[T any](
	f func(r Request) (T, errorx.HandledError),
) func(w http.ResponseWriter, r *http.Request) (T, errorx.HandledError) {
	return func(w http.ResponseWriter, r *http.Request) (T, errorx.HandledError) {
		req := Request{r}
		req.Logger()
		return f(Request{r})
	}
}
