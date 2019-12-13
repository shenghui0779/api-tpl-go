package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/iiinsomnia/demo/helpers"
	"github.com/iiinsomnia/demo/service"
)

func GetBookInfo(c *gin.Context) {
	s := new(service.BookInfo)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams), err.Error())

		return
	}

	resp, err := s.Do()

	if err != nil {
		Err(c, err)

		return
	}

	OK(c, resp)
}
