package helpers

import (
	"fmt"

	"github.com/shenghui0779/yiigo"
)

// Recover recover panic
func Recover() {
	if err := recover(); err != nil {
		yiigo.Logger().Error(fmt.Sprintf("pay-center panic: %v", err))
	}
}
