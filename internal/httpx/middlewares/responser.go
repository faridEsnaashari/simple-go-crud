package middlewares

import (
	"encoding/json"
	"net/http"
)

type ResponseStatus string

const (
	SuccessResponseStatus ResponseStatus = "success"
	FailedResponseStatus  ResponseStatus = "failed"
)

type Response struct {
	Status     ResponseStatus `json:"status"`
	Message    string         `json:"message"`
	HttpStatus int            `json:"httpStatus,omitempty"`
	Data       any            `json:"data"`
}

func WithResponser(f func(w http.ResponseWriter, r *http.Request) Response) func(w http.ResponseWriter, r *http.Request) Response {
	return func(w http.ResponseWriter, r *http.Request) Response {
		result := f(w, r)

		status := http.StatusOK
		if result.HttpStatus != 0 && result.HttpStatus != http.StatusOK {
			status = result.HttpStatus
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(result)

		return result
	}
}
