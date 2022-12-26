package result

import "net/http"

const CodeOK = 0

type Result interface {
	JSON(w http.ResponseWriter, r *http.Request)
}

type ResultOption func(r *response)

func Err(err error) ResultOption {
	return func(r *response) {
		r.Msg = err.Error()
	}
}

func Data(data interface{}) ResultOption {
	return func(r *response) {
		r.Data = data
	}
}

// New returns a new Result
func New(code int, msg string, options ...ResultOption) Result {
	resp := &response{
		Code: code,
		Msg:  msg,
	}

	if code != CodeOK {
		resp.Err = true
	}

	for _, f := range options {
		f(resp)
	}

	return resp
}
