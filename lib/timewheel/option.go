package timewheel

import (
	"context"
	"time"
)

// TWOption 时间轮选项
type TWOption func(tw *timewheel)

// WithErrLog 设置时间轮错误日志
func WithErrLog(fn func(ctx context.Context, v ...any)) TWOption {
	return func(tw *timewheel) {
		tw.log = fn
	}
}

// TKOption 时间轮任务选项
type TKOption func(t *Task)

// WithAttempts 指定任务重试次数；默认：1
func WithAttempts(attempts uint16) TKOption {
	return func(t *Task) {
		if attempts > 0 {
			t.maxAttempts = attempts
		}
	}
}

// WithDefer 指定任务延迟执行时间；默认：立即执行
func WithDefer(fn func(attempts uint16) time.Duration) TKOption {
	return func(t *Task) {
		if fn != nil {
			t.deferFn = fn
		}
	}
}
