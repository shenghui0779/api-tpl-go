package v1

import (
	"dao/mongo"
	"dao/mysql"
	"dao/redis"
	"service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	id := c.Query("id")

	if strings.Trim(id, " ") == "" {
		c.JSON(200, gin.H{
			"code":    1002,
			"message": "缺少参数",
		})

		return
	}

	_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    1002,
			"message": "无效的参数",
		})

		return
	}

	data, sysErr := service.GetAdminInfo(_id)

	if sysErr != nil {
		var msg string

		if sysErr.Error() == "record not found" {
			msg = "用户不存在"
		} else {
			msg = "服务器内部错误"
		}

		c.JSON(200, gin.H{
			"code":    1002,
			"message": msg,
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    1001,
		"message": "success",
		"data":    data,
	})
}

func TestMysql(c *gin.Context) {
	testMysql := mysql.NewAdminDao()

	err := testMysql.AddAdmin("Lucifer", "5e3141b16636dd68b5c61fcf04f6e258", "J!ai7!AnpmMU@GqO", 1, "末日使者", 1)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    1001,
			"message": "success",
		})
	}
}

func TestRedis(c *gin.Context) {
	testRedis := redis.NewTestRedis()
	key := "iiinsomnia"
	fields := []interface{}{1, 2, 3}
	data := []string{}
	testRedis.RedisBase.HMGet(&data, key, fields)

	c.JSON(200, gin.H{
		"code":    1001,
		"message": "success",
		"data":    data,
	})
}

func TestMongo(c *gin.Context) {
	testMongo := mongo.NewTestMongo()

	err := testMongo.AddPerson("yiiv587", 1, 20)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    1001,
			"message": "success",
		})
	}
}
