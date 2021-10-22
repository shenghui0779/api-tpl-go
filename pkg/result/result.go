package result

import (
	"net/http"
)

type Result interface {
	Wrap(options ...ResultOption) Result
	JSON(w http.ResponseWriter, r *http.Request)
}

type ResultOption func(r *response)

func WithErr(err error) ResultOption {
	return func(r *response) {
		r.err = err
	}
}

func WithData(data interface{}) ResultOption {
	return func(r *response) {
		r.data = data
	}
}

// New returns a new Result
func New(code int, msg string) Result {
	return &response{
		code: code,
		msg:  msg,
	}
}
