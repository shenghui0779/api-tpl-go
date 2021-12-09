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
	err  error
	data interface{}
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	obj := yiigo.X{
		"code": resp.code,
		"err":  false,
		"msg":  fmt.Sprintf("[%s] %s", logger.GetReqID(r.Context()), resp.err.Error()),
	}

	if resp.code != 0 {
		obj["err"] = true
	}

	if resp.data != nil {
		obj["data"] = resp.data
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
