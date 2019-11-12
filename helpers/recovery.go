package helpers

import (
	"fmt"

	"github.com/iiinsomnia/yiigo/v4"
)

// Recover recover panic
func Recover() {
	if err := recover(); err != nil {
		yiigo.Logger().Error(fmt.Sprintf("yiigo-demo panic: %v", err))
	}

	return
}
