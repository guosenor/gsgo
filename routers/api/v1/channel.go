package v1

import (
	"gsgo/models"
	"gsgo/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetChannels
// @Tags  频道
// @Summary get channels
// @Security ApiKeyAuth
// @Description get all channel
// @Produce  json
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /channels [get]
func GetChannels(c *gin.Context) {
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["channels"] = models.GetChannels()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
