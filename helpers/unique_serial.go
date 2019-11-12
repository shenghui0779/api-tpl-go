package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// UniqueSerial 生成唯一序列
func UniqueSerial() int64 {
	now := time.Now()

	r := rand.New(rand.NewSource(now.UnixNano()))
	s := fmt.Sprintf("%d%d", now.Nanosecond(), r.Intn(100))
	i, _ := strconv.ParseInt(s, 10, 64)

	return i
}
