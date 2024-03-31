package main

import (
	"api/cmd"
	"api/lib/db"
	"api/lib/redis"
)

func main() {
	defer clean()
	cmd.Init()
}

// clean 清理资源
func clean() {
	// 关闭数据库连接
	if cli := db.Client(); cli != nil {
		cli.Close()
	}
	// 关闭Redis连接
	if cli := redis.Client(); cli != nil {
		cli.Close()
	}
	// 关闭Redis集群连接
	if cli := redis.Cluster(); cli != nil {
		cli.Close()
	}
}
