package router

import (
	. "first_gin-gorm_proj/api/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/user/Login", Login)

	router.POST("/user/register", Register)

	router.POST("/user/updatepwd", Update)

	router.DELETE("/user/delete", Delete)

	return router
}
