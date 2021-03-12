package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"github.com/shenghui0779/demo/helpers"
)

var validator = yiigo.NewGinValidator()

func BindJSON(r *http.Request, obj interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return helpers.Err(helpers.ErrParams, err.Error())
	}

	if reflect.Indirect(reflect.ValueOf(obj)).Kind() != reflect.Struct {
		return nil
	}

	if err := validator.ValidateStruct(obj); err != nil {
		return helpers.Err(helpers.ErrParams, err.Error())
	}

	return nil
}

// OK returns success of an API.
func OK(w http.ResponseWriter, data ...interface{}) {
	obj := yiigo.X{
		"err":  false,
		"code": 0,
		"msg":  "success",
	}

	if len(data) > 0 {
		obj["data"] = data[0]
	}

	helpers.JSON(w, obj)
}

// Err returns error of an API.
func Err(w http.ResponseWriter, r *http.Request, err error) {
	code := helpers.ErrCode(err)
	msg := helpers.ErrMsg(err)

	if code == helpers.ErrSystem {
		yiigo.Logger().Error(fmt.Sprintf("server error: %d | %s", code, msg),
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
			zap.String("request_id", middleware.GetReqID(r.Context())),
			zap.Error(err),
		)
	}

	obj := yiigo.X{
		"err":  true,
		"code": code,
		"msg":  msg,
	}

	helpers.JSON(w, obj)
}
