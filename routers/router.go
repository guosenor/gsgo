package routers

import (
	"github.com/gin-gonic/gin"

	"gsgo/pkg/setting"
	v1 "gsgo/routers/api/v1"
)

// InitRouter init
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	apiV1 := r.Group("/api/v1")
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
	}

	return r
}
