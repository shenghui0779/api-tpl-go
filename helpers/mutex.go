package helpers

import (
	"time"

	"github.com/iiinsomnia/yiigo/v4"
	"go.uber.org/zap"
)

// Lock 加锁
// waitTime 等待时间，单位：毫秒
// expireTime 过期时间，单位：秒
func Lock(mutexKey string, waitTime int, expireTime int) {
	conn, err := yiigo.Redis().Get()

	if err != nil {
		yiigo.Logger().Error("mutex lock error", zap.Error(err))

		return
	}

	defer yiigo.Redis().Put(conn)

	maxLoop := expireTime * 1000 / waitTime

	for i := 0; i < maxLoop; i++ {
		ok, err := conn.Do("SET", mutexKey, time.Now().Nanosecond(), "EX", expireTime, "NX")

		if err != nil {
			yiigo.Logger().Error("mutex lock error", zap.Error(err))

			return
		}

		if ok != nil {
			return
		}

		time.Sleep(time.Duration(waitTime) * time.Millisecond)
	}
}

// UnLock 解锁
func UnLock(mutexKey string) {
	conn, err := yiigo.Redis().Get()

	if err != nil {
		yiigo.Logger().Error("mutex unlock error", zap.Error(err))

		return
	}

	defer yiigo.Redis().Put(conn)

	_, err = conn.Do("DEL", mutexKey)

	if err != nil {
		yiigo.Logger().Error("mutex unlock error", zap.Error(err))
	}
}
