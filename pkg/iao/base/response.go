package base

type Response interface {
	OK() bool
	ErrCode() int
	ErrMsg() string
}

type response struct {
	Code int         `json:"code"`
	Err  bool        `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *response) OK() bool {
	return !r.Err
}

func (r *response) ErrCode() int {
	return r.Code
}

func (r *response) ErrMsg() string {
	return r.Msg
}

// NewResponse returns new response.
// Note: param data should be a pointer.
func NewResponse(data ...interface{}) Response {
	if len(data) == 0 {
		return new(response)
	}

	return &response{Data: data[0]}
}
