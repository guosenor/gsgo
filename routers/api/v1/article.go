package v1

import (
	"fmt"
	"gsgo/models"
	"gsgo/pkg/e"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetArticleByID id
// @Tags  文章
// @Summary get a article by id
// @Description get article by ID
// @ID tagId
// @Accept  json
// @Produce  json
// @Param id path int true "aticle ID"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /articles/{id} [get]
func GetArticleByID(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	code := e.SUCCESS
	article := models.GetArticleByID(id)
	if article.ID != 0 {
		maps["ID"] = id
		data["atricle"] = article
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type createArticle struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ChannelID int    `json:"channelId"`
}

// AddArticle article
// @Summary AddArticle
// @Security ApiKeyAuth
// @Produce  json
// @Tags  文章
// @Param body body v1.createArticle true "新建"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /articles [post]
func AddArticle(c *gin.Context) {
	userId, _ := c.MustGet("userId").(int)

	fmt.Println(" auth userId:", userId)
	var article createArticle
	c.BindJSON(&article)
	valid := validation.Validation{}
	valid.Required(article.Title, "title").Message("标题不能为空")
	valid.MaxSize(article.Title, 100, "title").Message("标题100字符")
	valid.MinSize(article.Title, 2, "title").Message("标题不能少于2字符")
	valid.Required(article.Body, "body").Message("标题不能为空")
	valid.MaxSize(article.Body, 1000, "body").Message("标题100字符")
	valid.MinSize(article.Body, 2, "body").Message("标题不能少于2字符")

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = e.SUCCESS
		maps := make(map[string]interface{})
		maps["title"] = article.Title
		data["article"] = models.AddArticle(article.Title, article.Body, article.ChannelID, userId)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
