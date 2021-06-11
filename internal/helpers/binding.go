package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"tplgo/internal/result"

	"github.com/shenghui0779/yiigo"
)

var validator = yiigo.NewValidator()

func BindJSON(r *http.Request, obj interface{}) result.Result {
	defer io.Copy(io.Discard, r.Body)

	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return result.ErrParams.Wrap(result.WithMsg(err.Error()))
	}

	if err := validator.ValidateStruct(obj); err != nil {
		return result.ErrParams.Wrap(result.WithMsg(err.Error()))
	}

	return nil
}
