package result

import (
	"api/logger"
	"encoding/json"
	"net/http"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

const CodeOK = 0

type Result interface {
	JSON(w http.ResponseWriter, r *http.Request)
}

type response struct {
	x yiigo.X
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp.x["req_id"] = logger.GetReqID(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp.x); err != nil {
		logger.Err(ctx, "err write response", zap.Error(err))
	}
}

type Option func(r *response)

// M 指定返回的Msg
func M(m string) Option {
	return func(r *response) {
		r.x["msg"] = m
	}
}

// V 指定返回的Data
func V(v any) Option {
	return func(r *response) {
		r.x["data"] = v
	}
}

// KV 指定返回的自定义key-value
func KV(k string, v any) Option {
	return func(r *response) {
		r.x[k] = v
	}
}

// New returns a new Result
func New(code int, msg string, options ...Option) Result {
	resp := &response{
		x: yiigo.X{
			"code": code,
			"err":  false,
			"msg":  msg,
		},
	}

	if code != CodeOK {
		resp.x["err"] = true
	}

	for _, f := range options {
		f(resp)
	}

	return resp
}
