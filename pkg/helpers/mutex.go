package helpers

import (
	"context"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gomodule/redigo/redis"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/pkg/result"
)

type MutexProcessFunc func(ctx context.Context) result.Result

// Mutex is a reader/writer mutual exclusion lock.
type Mutex interface {
	// Acquire 获取锁
	// interval 每次获取锁的间隔时间（毫秒）
	Acquire(ctx context.Context, process MutexProcessFunc, timeout time.Duration) result.Result
}

type distributed struct {
	key     string
	timeout time.Duration
}

func (d *distributed) Acquire(ctx context.Context, process MutexProcessFunc, interval time.Duration) result.Result {
	if deadline, ok := ctx.Deadline(); ok {
		if v := time.Until(deadline); v < d.timeout {
			d.timeout = v
		}
	}

	conn, err := yiigo.Redis().Get(ctx)

	// 获取锁出错
	if err != nil {
		yiigo.Logger().Error("[mutex] acquire lock error", zap.Error(err), zap.String("request_id", middleware.GetReqID(ctx)))

		return result.ErrSystem
	}

	defer yiigo.Redis().Put(conn)

	// 缓存超时时间
	ex := int64(d.timeout.Seconds())

	for {
		select {
		case <-ctx.Done():
			// 锁获取超时或被取消
			yiigo.Logger().Warn("[mutex] acquire lock error", zap.Error(ctx.Err()), zap.String("request_id", middleware.GetReqID(ctx)), zap.Duration("timeout", d.timeout))

			return result.ErrTimeout
		default:
		}

		// 获取锁
		reply, err := redis.String(conn.Do("SET", d.key, time.Now().Nanosecond(), "EX", ex, "NX"))

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
func DistributedMutex(key string, timeout time.Duration) Mutex {
	mutex := &distributed{
		key:      key,
		timeout:  10 * time.Second,
	}

	if timeout > 0 {
		mutex.timeout = timeout
	}

	return mutex
}
