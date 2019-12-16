package v1

import (
	"fmt"
	"gsgo/models"
	"gsgo/pkg/e"
	"gsgo/service/article_service"
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
	articleService := article_service.Article{ID: id}
	article := articleService.GetCache()
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
		data["article"] = models.AddArticle(article.Title, article.Body, article.ChannelID, userId)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle article
// @Summary AddArticle
// @Security ApiKeyAuth
// @Produce  json
// @Tags  文章
// @Param id path int true "aticle ID"
// @Param body body v1.createArticle true "新建"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /articles/{id} [put]
func UpdateArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	userId, _ := c.MustGet("userId").(int)
	form := make(map[string]interface{})
	data := make(map[string]interface{})
	var article createArticle
	c.BindJSON(&article)
	valid := validation.Validation{}
	if article.Title != "" {
		form["Title"] = article.Title
		valid.MaxSize(article.Title, 16, "title").Message("标题不能少于2字符")
	}
	if article.Body != "" {
		form["Body"] = article.Body
		valid.MaxSize(article.Body, 1000, "body").Message("标题100字符")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["article"] = models.UpdateArticleById(id, userId, form)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles multiple
// @Tags  文章
// @Summary get multiple article
// @Description get multiple  article
// @Accept  json
// @Produce  json
// @Param channelId query int false "channel ID"
// @Param createById query int false "createBy ID"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	channelId := com.StrTo(c.Query("channelId")).MustInt()
	createById := com.StrTo(c.Query("createById")).MustInt()
	article := article_service.Article{}
	if channelId != 0 {
		article.ChannelID = channelId
	}
	if channelId != 0 {
		article.CreateByID = createById
	}
	data := make(map[string]interface{})
	data["Articles"] = article.GetArticlesCache()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
