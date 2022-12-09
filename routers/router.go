package routers

import (
	"ChsSeltApi/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	index := router.Group("/")
	{
		index.GET("/login", controller.Login.Login)
		index.POST("/login", controller.Login.Login)
	}
	router.Run(":8088")
}
