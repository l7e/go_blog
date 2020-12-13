package routers

import (
	"github.com/gin-gonic/gin"
	setting "go_blog/pkg"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello",
		})
	})

	return r
}
