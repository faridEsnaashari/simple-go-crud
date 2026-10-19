package middlewares

import (
	"net/http"
	"time"

	logx "a.com/test/internal/httpx/logs"
)

type Request struct {
	*http.Request
}

func (r *Request) Logger(status int, duration int64) {
	logx.Info("INCOMMING_REQ", "method", r.Method, "route", r.Pattern, "path", r.URL.Path, "status", status, "duration", duration)
}

func WithLogger(
	f func(w http.ResponseWriter, r *http.Request) Response,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		req := Request{r}
		result := f(w, r)

		status := http.StatusOK
		if result.HttpStatus != 0 {
			status = result.HttpStatus
		}

		duration := time.Since(begin).Milliseconds()

		req.Logger(status, duration)
	}
}
