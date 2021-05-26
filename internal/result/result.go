package result

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type Result interface {
	Wrap(options ...ResultOption) Result
	JSON(w http.ResponseWriter, r *http.Request)
}

type response struct {
	code int
	msg  string
	err  error
	data interface{}
}

func (resp *response) Wrap(options ...ResultOption) Result {
	newR := resp.clone()

	if len(options) != 0 {
		for _, f := range options {
			f(newR)
		}
	}

	return newR
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	obj := yiigo.X{
		"code": resp.code,
		"err":  false,
		"msg":  resp.msg,
	}

	if resp.code != 0 {
		obj["err"] = true
	}

	if resp.data != nil {
		obj["data"] = resp.data
	}

	// 如有错误，记录日志
	if resp.err != nil {
		yiigo.Logger().Error(fmt.Sprintf("Whoops! Server Error: %d | %s", resp.code, resp.msg),
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
			zap.String("request_id", middleware.GetReqID(r.Context())),
			zap.Error(resp.err),
		)
	}

	b, err := json.Marshal(obj)

	if err != nil {
		yiigo.Logger().Error("Response JSON error", zap.Error(err))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if _, err = w.Write(b); err != nil {
		yiigo.Logger().Error("Response JSON error", zap.Error(err))
	}
}

func (resp *response) clone() *response {
	return &response{
		code: resp.code,
		msg:  resp.msg,
		err:  resp.err,
		data: resp.data,
	}
}

// New returns a new Result
func New(code int, msg string) Result {
	return &response{
		code: code,
		msg:  msg,
	}
}
