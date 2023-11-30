package nsq

import (
	"api/lib/log"

	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

var producer *nsq.Producer

// Logger 实现nsq日志接口
type Logger struct{}

// Output nsq错误输出
func (l *Logger) Output(calldepth int, s string) error {
	log.Error(context.Background(), fmt.Sprintf("err nsq: %s", s), zap.Int("call_depth", calldepth))

	return nil
}

func Init(nsqd string, lookupd []string, cfg *nsq.Config, consumers ...Consumer) error {
	if cfg == nil {
		cfg = nsq.NewConfig()
	}

	var err error

	producer, err = nsq.NewProducer(nsqd, cfg)
	if err != nil {
		return err
	}
	if err = producer.Ping(); err != nil {
		return err
	}

	producer.SetLogger(&Logger{}, nsq.LogLevelError)

	// set consumers
	if err = setConsumers(lookupd, consumers...); err != nil {
		return err
	}

	return nil
}

// Publish 同步推送消息到指定Topic
func Publish(topic string, msg []byte) error {
	if producer == nil {
		return errors.New("nsq producer is nil (forgotten init?)")
	}

	return producer.Publish(topic, msg)
}

// DeferredPublish 同步推送延迟消息到指定Topic
func DeferredPublish(topic string, msg []byte, duration time.Duration) error {
	if producer == nil {
		return errors.New("nsq producer is nil (forgotten init?)")
	}

	return producer.DeferredPublish(topic, duration, msg)
}

// NextAttemptDelay 一个帮助方法，用于返回下一次尝试的等待时间
func NextAttemptDelay(attempts uint16) time.Duration {
	var d time.Duration

	switch attempts {
	case 0, 1:
		d = 5 * time.Second
	case 2:
		d = 10 * time.Second
	case 3:
		d = 15 * time.Second
	case 4:
		d = 30 * time.Second
	case 5:
		d = time.Minute
	case 6:
		d = 2 * time.Minute
	case 7:
		d = 5 * time.Minute
	case 8:
		d = 10 * time.Minute
	case 9:
		d = 15 * time.Minute
	case 10:
		d = 30 * time.Minute
	default:
		d = time.Hour
	}

	return d
}
