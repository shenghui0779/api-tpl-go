package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo_demo/helpers"
	"github.com/iiinsomnia/yiigo_demo/service"
	"github.com/pkg/errors"
)

func GetBookInfo(c *gin.Context) {
	s := new(service.BookInfo)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book info params error")))

		return
	}

	if _, err := govalidator.ValidateStruct(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book info params error")), err.Error())

		return
	}

	resp, err := s.Do()

	if err != nil {
		Err(c, err)

		return
	}

	OK(c, resp)
}

func AddBook(c *gin.Context) {
	s := new(service.BookAdd)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book add params error")))

		return
	}

	if _, err := govalidator.ValidateStruct(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book add params error")), err.Error())

		return
	}

	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

func EditBook(c *gin.Context) {
	s := new(service.BookEdit)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book edit params error")))

		return
	}

	if _, err := govalidator.ValidateStruct(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book edit params error")), err.Error())

		return
	}

	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

func DeleteBook(c *gin.Context) {
	s := new(service.BookDelete)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book delete params error")))

		return
	}

	if _, err := govalidator.ValidateStruct(s); err != nil {
		Err(c, helpers.Error(10000, errors.Wrap(err, "book delete params error")), err.Error())

		return
	}

	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}
