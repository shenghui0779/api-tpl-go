package helpers

import (
	"context"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Lock 加锁
// waitTime 等待时间，单位：毫秒
// expireTime 过期时间，单位：秒
func Lock(mutexKey string, waitTime int, expireTime int) {
	conn, err := yiigo.Redis().Get(context.Background())

	if err != nil {
		yiigo.Logger().Error("mutex lock error", zap.Error(err))

		return
	}

	defer yiigo.Redis().Put(conn)

	duration := expireTime * 1000

	for {
		if duration <= 0 {
			return
		}

		if ok, err := conn.Do("SET", mutexKey, time.Now().Nanosecond(), "EX", expireTime, "NX"); err != nil || ok != nil {
			if err != nil {
				yiigo.Logger().Error("mutex lock error", zap.Error(err))
			}

			return
		}

		time.Sleep(time.Duration(waitTime) * time.Millisecond)

		duration -= waitTime
	}
}

// UnLock 解锁
func UnLock(mutexKey string) {
	conn, err := yiigo.Redis().Get(context.Background())

	if err != nil {
		yiigo.Logger().Error("mutex unlock error", zap.Error(err))

		return
	}

	defer yiigo.Redis().Put(conn)

	if _, err = conn.Do("DEL", mutexKey); err != nil {
		yiigo.Logger().Error("mutex unlock error", zap.Error(err))
	}
}
