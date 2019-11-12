package helpers

import (
	"fmt"

	"github.com/iiinsomnia/yiigo/v4"
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

	fields := make([]zap.Field, 0, 1)

	if len(err) > 0 {
		fields = append(fields, zap.Error(err[0]))
	}

	yiigo.Logger().Error(fmt.Sprintf("Whoops! %d | %s", code, msg), fields...)

	return &Err{
		code: code,
		msg:  msg,
	}
}

// ErrNoLog returns an error
func ErrNoLog(code int) error {
	msg := "unkown code"

	if m, ok := errCodes[code]; ok {
		msg = m
	}

	return &Err{
		code: code,
		msg:  msg,
	}
}
