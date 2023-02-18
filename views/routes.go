package views

import (
	"user_proj/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/locations/:id", controller.Example)
	return r
}
