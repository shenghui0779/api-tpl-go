package result

import (
	"encoding/json"
	"net/http"

	"api/pkg/logger"

	"go.uber.org/zap"
)

type response struct {
	Code  int    `json:"code"`
	Err   bool   `json:"err"`
	Msg   string `json:"msg"`
	ReqID string `json:"req_id"`
	Data  any    `json:"data,omitempty"`
}

func (resp *response) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp.ReqID = logger.GetReqID(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Err(ctx, "err write response", zap.Error(err))
	}
}
