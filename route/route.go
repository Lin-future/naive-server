package route

import (
	"naive-server/controller"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetupRouter() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	r.POST("/checkin", controller.CheckIn)
	r.Run(":8080")
}
