package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var cli redis.UniversalClient

// Init 初始化Redis
func Init(opts *redis.UniversalOptions) error {
	cli = redis.NewUniversalClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// verify connection
	if err := cli.Ping(ctx).Err(); err != nil {
		cli.Close()
		return err
	}

	return nil
}

func Client() redis.UniversalClient {
	return cli
}
