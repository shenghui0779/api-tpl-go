package helpers

import (
	"context"
	"time"
	"tplgo/internal/result"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gomodule/redigo/redis"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Mutex 分布式互斥锁
// interval：每次获取锁的等待时间（毫秒）
// timeout：锁获取超时时间（秒）
func Mutex(ctx context.Context, key string, process func() result.Result, interval, timeout time.Duration) result.Result {
	conn, err := yiigo.Redis().Get(ctx)

	// 获取锁出错，直接执行任务
	if err != nil {
		yiigo.Logger().Error("acquire mutex error", zap.Error(err), zap.String("request_id", middleware.GetReqID(ctx)))

		return process()
	}

	defer yiigo.Redis().Put(conn)

	now := time.Now().Local()
	expire := int64(timeout.Seconds())

	for {
		// 锁获取超时，执行任务
		if time.Since(now).Seconds() >= timeout.Seconds() {
			yiigo.Logger().Warn("acquire mutex timeout", zap.String("timeout", timeout.String()), zap.String("request_id", middleware.GetReqID(ctx)))

			return process()
		}

		// 获取锁
		reply, err := redis.String(conn.Do("SET", key, time.Now().Nanosecond(), "EX", expire, "NX"))

		// 获取锁出错，直接执行任务
		if err != nil && err != redis.ErrNil {
			yiigo.Logger().Error("acquire mutex error", zap.Error(err), zap.String("request_id", middleware.GetReqID(ctx)))

			return process()
		}

		// 获取锁成功，结束等待，执行任务
		if reply == "OK" {
			break
		}

		// 未能获取到锁，等待下一次获取
		time.Sleep(interval)
	}

	// 任务执行结束，释放锁
	defer conn.Do("DEL", key)

	return process()
}
