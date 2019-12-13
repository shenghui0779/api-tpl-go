package helpers

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/iiinsomnia/yiigo/v4"
)

var BufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) // 16KB
	},
}

// Recover recover panic
func Recover() {
	if err := recover(); err != nil {
		yiigo.Logger().Error(fmt.Sprintf("pay-center panic: %v", err))
	}

	return
}
