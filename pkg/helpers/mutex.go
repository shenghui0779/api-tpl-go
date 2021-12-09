package helpers

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
)

type MutexHandler func(ctx context.Context) error

// Mutex is a reader/writer mutual exclusion lock.
type Mutex interface {
	// Acquire 获取锁
	// interval 每次获取锁的间隔时间（每隔interval时间尝试获取一次锁）
	// timeout 锁获取超时时间
	Acquire(ctx context.Context, callback MutexHandler, interval, timeout time.Duration) error
}

type distributed struct {
	key    string
	expire int64
}

func (d *distributed) Acquire(ctx context.Context, callback MutexHandler, interval, timeout time.Duration) error {
	mutexCtx := ctx

	if timeout > 0 {
		var cancel context.CancelFunc

		mutexCtx, cancel = context.WithTimeout(mutexCtx, timeout)

		defer cancel()
	}

	conn, err := yiigo.Redis().Get(mutexCtx)

	// 获取锁出错
	if err != nil {
		return errors.Wrap(err, "err redis conn")
	}

	defer yiigo.Redis().Put(conn)

	for {
		select {
		case <-mutexCtx.Done():
			// 锁获取超时或被取消
			return errors.Wrap(mutexCtx.Err(), "err mutex context")
		default:
		}

		// 获取锁
		reply, err := redis.String(conn.Do("SET", d.key, time.Now().Nanosecond(), "EX", d.expire, "NX"))

		// 获取锁出错
		if err != nil && err != redis.ErrNil {
			return errors.Wrap(err, "err redis setnx")
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

	return callback(ctx)
}

// DistributedMutex returns is a distributed mutual exclusion lock.
func DistributedMutex(key string, expire time.Duration) Mutex {
	mutex := &distributed{
		key:    key,
		expire: 10,
	}

	if seconds := expire.Seconds(); seconds > 0 {
		mutex.expire = int64(seconds)
	}

	return mutex
}
