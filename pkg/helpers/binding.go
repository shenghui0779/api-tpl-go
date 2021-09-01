package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/yiigo"

	"tplgo/pkg/result"
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

func URLParamInt(r *http.Request, key string) (int64, result.Result) {
	v, err := strconv.ParseInt(chi.URLParam(r, key), 10, 64)

	if err != nil {
		return 0, result.ErrParams.Wrap(result.WithMsg(err.Error()))
	}

	return v, nil
}
