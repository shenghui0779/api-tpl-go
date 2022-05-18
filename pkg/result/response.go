package result

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tplgo/pkg/logger"

	"go.uber.org/zap"
)

type response struct {
	Code  int         `json:"code"`
	Err   bool        `json:"err"`
	Msg   string      `json:"msg"`
	ReqID string      `json:"req_id"`
	Data  interface{} `json:"data,omitempty"`
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp.ReqID = logger.GetReqID(ctx)

	if resp.Code != CodeOK {
		resp.Err = true
	}

	b, err := json.Marshal(resp)

	if err != nil {
		logger.Err(ctx, "err marshal response to JSON", zap.Error(err))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("[%s] %s", logger.GetReqID(ctx), err.Error())))

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(b); err != nil {
		logger.Err(ctx, "err write response", zap.Error(err))
	}
}
