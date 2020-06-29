package helpers

import (
	"fmt"

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
func Error(code int, err ...error) error {
	msg := "unkown code"

	if m, ok := errCodes[code]; ok {
		msg = m
	}

	if len(err) > 0 {
		yiigo.Logger().Error(fmt.Sprintf("Whoops! %d | %s", code, msg), zap.Error(err[0]))
	}

	return &Err{
		code: code,
		msg:  msg,
	}
}
