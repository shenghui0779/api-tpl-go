package routes

import (
	"controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitTestRoutes(router *gin.Engine) {
	//对内
	in := router.Group("/i")
	{
		inV1 := in.Group("/v1")
		{
			inV1.GET("test", v1.Test)
			inV1.GET("test/mysql", v1.TestMysql)
			inV1.GET("test/redis", v1.TestRedis)
			inV1.GET("test/mongo", v1.TestMongo)
		}
	}

	//对外
	out := router.Group("/api")
	{
		outV1 := out.Group("/v1")
		{
			outV1.GET("test", v1.Test)
			outV1.GET("test/mysql", v1.TestMysql)
			outV1.GET("test/redis", v1.TestRedis)
			outV1.GET("test/mongo", v1.TestMongo)
		}
	}
}
