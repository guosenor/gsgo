package routers

import (
	"github.com/gin-gonic/gin"

	"gsgo/middleware/jwt"
	"gsgo/pkg/setting"
	api "gsgo/routers/api"
	v1 "gsgo/routers/api/v1"
)

// InitRouter init
func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.POST("/api/v1/auth", api.GetAuth)
	apiV1 := r.Group("/api/v1")
	// 获取指定文章
	apiV1.GET("/articles/:id", v1.GetArticleByID)
	// 获取全部频道
	apiV1.GET("/channels", v1.GetChannels)
	apiV1.Use(jwt.JWT())
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//获取标签
		apiV1.GET("/tags/:id", v1.GetTagByID)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DelTagByID)
		// 新建文章
		apiV1.POST("/articles", v1.AddArticle)
		// 修改文章
		apiV1.PUT("/articles/:id", v1.UpdateArticle)
	}
	return r
}
