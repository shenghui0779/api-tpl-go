package helpers

import (
	"fmt"

	"gitlab.meipian.cn/golib/yiigo"
)

// Recover recover panic
func Recover() {
	if err := recover(); err != nil {
		yiigo.Logger().Error(fmt.Sprintf("yiigo_demo panic: %v", err))
	}

	return
}
