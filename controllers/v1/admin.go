package v1

import (
	"service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAdminDetail(c *gin.Context) {
	id := c.Query("id")

	if strings.TrimSpace(id) == "" {
		c.JSON(200, gin.H{"code": -1, "msg": "参数错误"})
		return
	}

	_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": "参数错误"})
		return
	}

	data, err := service.GetAdminById(_id)

	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": "服务器内部错误"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "请求成功", "data": data})
}

func GetAdminList(c *gin.Context) {
	data, count, err := service.GetAdminList()

	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": "服务器内部错误"})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "请求成功",
		"data": map[string]interface{}{
			"total": count,
			"list":  data,
		},
	})
}
