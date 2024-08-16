package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"api/app/ent"
	"api/lib/log"
	"api/lib/redis"
)

func TestMain(m *testing.M) {
	// 加载配置(注意：替换成自己的配置文件路径)
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 初始化日志
	log.Init()
	// 初始化数据库
	if err := ent.Init(); err != nil {
		panic(err)
	}
	defer func() {
		_ = ent.DB.Close()
	}()
	// 初始化Redis
	if err := redis.Init(); err != nil {
		panic(err)
	}
	defer func() {
		_ = redis.Client.Close()
	}()
	m.Run()
}

func TestAuth(t *testing.T) {
	req := &ReqLogin{
		Account:  "demo",
		Password: "password",
	}
	ret := Login(context.Background(), req)
	fmt.Println("[Result] ---", ret)
}
