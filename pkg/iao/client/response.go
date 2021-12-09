package client

import "fmt"

type Response interface {
	Error() error
}

type response struct {
	Code int         `json:"code"`
	Err  bool        `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *response) Error() error {
	if !r.Err {
		return fmt.Errorf("%d|%s", r.Code, r.Msg)
	}

	return nil
}

// NewResponse returns new response.
// Note: param data should be a pointer.
func NewResponse(data ...interface{}) Response {
	if len(data) == 0 {
		return new(response)
	}

	return &response{Data: data[0]}
}
