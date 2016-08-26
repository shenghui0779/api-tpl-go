package redis

import "github.com/iiinsomnia/yiigo"

type TestRedis struct {
	yiigo.RedisBase
}

func NewTestRedis() *TestRedis {
	return &TestRedis{yiigo.RedisBase{CacheName: "test"}}
}
