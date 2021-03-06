package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/middleware/jwt"
	setting "go_blog/pkg"
	"go_blog/routers/api"
	v1 "go_blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) //使用中间件
	{
		apiv1.GET("/articles", v1.GetArticles)

		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTags)

		apiv1.PUT("/tags/:id", v1.EditTags)

		apiv1.DELETE("/tags/:id", v1.DeleteTags)
	}

	return r
}
