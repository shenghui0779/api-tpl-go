package v1

import (
	"service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func GetUserDetail(c *gin.Context) {
	id := c.Query("id")

	if strings.TrimSpace(id) == "" {
		yiigo.ReturnJson(c, -1, "参数错误")
		return
	}

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "参数错误")
		return
	}

	data, err := service.GetUserById(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "服务器内部错误")
		return
	}

	yiigo.ReturnJson(c, 0, "请求成功", data)
}

func GetUserList(c *gin.Context) {
	data, count, err := service.GetUserList()

	if err != nil {
		yiigo.ReturnJson(c, -1, "服务器内部错误")
		return
	}

	result := map[string]interface{}{
		"total": count,
		"list":  data,
	}

	yiigo.ReturnJson(c, 0, "请求成功", result)
}
