package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type StatusErr interface {
	Code() int
	Error() string
}

type Err struct {
	code int
	msg  string
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Error() string {
	return e.msg
}

// Error returns an error
func Error(ctx *gin.Context, code int, err ...error) error {
	msg := "unkown code"

	if m, ok := errCodes[code]; ok {
		msg = m
	}

	if len(err) > 0 {
		yiigo.Logger().Error(fmt.Sprintf("[%v] %d | %s", ctx.Request.URL, code, msg),
			zap.String("request_id", ctx.GetHeader("request_id")),
			zap.Error(err[0]),
		)
	}

	return &Err{
		code: code,
		msg:  msg,
	}
}
