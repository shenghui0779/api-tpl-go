package main

import (
	"api/app/cmd"
	"api/app/ent"
	"api/lib/redis"
)

func main() {
	defer clean()
	cmd.Init()
}

// clean 清理资源
func clean() {
	// 关闭数据库连接
	if ent.DB != nil {
		ent.DB.Close()
	}
	// 关闭Redis连接
	if redis.Client != nil {
		redis.Client.Close()
	}
	// 关闭Redis集群连接
	if redis.Cluster != nil {
		redis.Cluster.Close()
	}
}
