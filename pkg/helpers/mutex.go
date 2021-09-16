package helpers

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"

	"tplgo/pkg/result"
)

type MutexProcessFunc func(ctx context.Context) result.Result

// Mutex is a reader/writer mutual exclusion lock.
type Mutex interface {
	// Acquire 获取锁
	// interval 每次获取锁的间隔时间
	// expire 锁超时时间
	Acquire(ctx context.Context, process MutexProcessFunc, interval, expire time.Duration) result.Result
}

type distributed struct {
	key     string
	timeout time.Duration
}

func (d *distributed) Acquire(ctx context.Context, process MutexProcessFunc, interval, expire time.Duration) result.Result {
	mutexCtx := ctx

	if d.timeout > 0 {
		var cancel context.CancelFunc

		mutexCtx, cancel = context.WithTimeout(mutexCtx, d.timeout)

		defer cancel()
	}

	conn, err := yiigo.Redis().Get(mutexCtx)

	// 获取锁出错
	if err != nil {
		return result.ErrMutex.Wrap(result.WithErr(errors.Wrap(err, "acquire lock error")))
	}

	defer yiigo.Redis().Put(conn)

	// 缓存超时时间
	ex := int64(expire.Seconds())

	for {
		select {
		case <-mutexCtx.Done():
			// 锁获取超时或被取消
			return result.ErrMutex.Wrap(result.WithErr(errors.Wrap(mutexCtx.Err(), "acquire lock error")))
		default:
		}

		// 获取锁
		reply, err := redis.String(conn.Do("SET", d.key, time.Now().Nanosecond(), "EX", ex, "NX"))

		// 获取锁出错
		if err != nil && err != redis.ErrNil {
			return result.ErrMutex.Wrap(result.WithErr(errors.Wrap(err, "acquire lock error")))
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
	return &distributed{
		key:     key,
		timeout: timeout,
	}
}
