package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shenghui0779/demo/helpers"
)

// OK returns success of an API.
func OK(ctx *gin.Context, data ...interface{}) {
	obj := gin.H{
		"err":  false,
		"code": 1000,
		"msg":  "success",
	}

	if len(data) > 0 {
		obj["data"] = data[0]
	}

	ctx.Set("response", obj)

	ctx.JSON(http.StatusOK, obj)
}

// Err returns error of an API.
func Err(ctx *gin.Context, err error, msg ...string) {
	obj := gin.H{
		"err":  true,
		"code": 50000,
		"msg":  "服务器错误，请稍后重试",
	}

	if e, ok := err.(helpers.StatusErr); ok {
		obj["code"] = e.Code()
		obj["msg"] = e.Error()
	}

	if len(msg) > 0 {
		obj["msg"] = msg[0]
	}

	ctx.Set("response", obj)

	ctx.JSON(http.StatusOK, obj)
}
