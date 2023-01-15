package routes

import (
	"bluebell/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成release模式
	}
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
