package helpers

import (
	"context"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gomodule/redigo/redis"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/internal/result"
)

// Mutex is a reader/writer mutual exclusion lock.
type Mutex interface {
	// Acquire 获取锁
	// interval 每次获取锁的间隔时间（毫秒）
	Acquire(ctx context.Context, process func(ctx context.Context) result.Result, interval time.Duration) result.Result
}

type distributed struct {
	key    string
	expire int64
}

func (d *distributed) Acquire(ctx context.Context, process func(ctx context.Context) result.Result, interval time.Duration) result.Result {
	conn, err := yiigo.Redis().Get(ctx)

	// 获取锁出错
	if err != nil {
		yiigo.Logger().Error("[mutex] acquire lock error", zap.Error(err), zap.String("request_id", middleware.GetReqID(ctx)))

		return result.ErrSystem
	}

	defer yiigo.Redis().Put(conn)

	for {
		select {
		case <-ctx.Done():
			// 锁获取超时
			yiigo.Logger().Warn("[mutex] acquire lock error", zap.Error(ctx.Err()), zap.String("request_id", middleware.GetReqID(ctx)))

			return result.ErrTimeout
		default:
		}

		// 获取锁
		reply, err := redis.String(conn.Do("SET", d.key, time.Now().Nanosecond(), "EX", d.expire, "NX"))

		// 获取锁出错
		if err != nil && err != redis.ErrNil {
			yiigo.Logger().Error("[mutex] acquire lock error", zap.Error(err), zap.String("request_id", middleware.GetReqID(ctx)))

			return result.ErrSystem
		}

		// 获取锁成功，结束等待，执行任务
		if reply == "OK" {
			break
		}

		// 未能获取到锁，等待下一次获取
		time.Sleep(interval)
	}

	// 任务执行结束，释放锁
	defer conn.Do("DEL", d.key)
	defer Recover(ctx)

	return process(ctx)
}

// DistributedMutex returns is a distributed mutual exclusion lock.
func DistributedMutex(key string, expire time.Duration) Mutex {
	mutex := &distributed{
		key:    key,
		expire: 10,
	}

	if v := int64(expire.Seconds()); v != 0 {
		mutex.expire = v
	}

	return mutex
}
