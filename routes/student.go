package routes

import (
	"demo/controllers/v1"

	"github.com/gin-gonic/gin"
)

func LoadStudentRoutes(router *gin.Engine) {
	_v1 := router.Group("/v1")
	// _v1.Use(middlewares.Auth())
	{
		_v1.GET("students", v1.GetStudentList)
		_v1.GET("students/:id", v1.GetStudentDetail)
		_v1.POST("students", v1.AddNewStudent)
		_v1.PUT("students/:id", v1.UpdateStudent)
		_v1.DELETE("students/:id", v1.DeleteStudent)
	}
}
