package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/shenghui0779/demo/helpers"
	"github.com/shenghui0779/demo/service"
)

func GetBookInfo(ctx *gin.Context) {
	s := new(service.BookInfo)

	if err := ctx.ShouldBindJSON(s); err != nil {
		Err(ctx, helpers.Error(ctx, helpers.ErrParams), err.Error())

		return
	}

	resp, err := s.Do(ctx)

	if err != nil {
		Err(ctx, err)

		return
	}

	OK(ctx, resp)
}
