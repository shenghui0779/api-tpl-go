package result

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tplgo/pkg/logger"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

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

	// 如有error，返回error信息
	if resp.err != nil {
		obj["msg"] = fmt.Sprintf("[%s] %s", logger.GetReqID(r.Context()), resp.err.Error())
	}

	b, err := json.Marshal(obj)

	if err != nil {
		logger.Err(r.Context(), "Response JSON error", zap.Error(err))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("[%s] %s", logger.GetReqID(r.Context()), err.Error())))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if _, err = w.Write(b); err != nil {
		logger.Err(r.Context(), "Response JSON error", zap.Error(err))
	}
}

func (resp *response) clone() *response {
	newResp := new(response)

	newResp.code = resp.code
	newResp.msg = resp.msg
	newResp.err = resp.err
	newResp.data = resp.data

	return newResp
}
