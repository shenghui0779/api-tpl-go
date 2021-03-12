package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func JSON(w http.ResponseWriter, obj yiigo.X) {
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
