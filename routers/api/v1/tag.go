package v1

import (
	"fmt"
	"gsgo/models"
	"gsgo/pkg/e"
	"gsgo/pkg/setting"
	"gsgo/pkg/util"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetTags get
// @Summary get tags list
// @Tags  标签
// @Description get tags list
// @Accept  json
// @Param name query string false "name"
// @Param page query int true "page"
// @Produce  json
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Header 200 {string} Token "qwerty"
// @Router /tags/ [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type tagCreate struct {
	Name  string `json:"name" example:"account name"`
	State int    `json:"state" example:"0"`
}

// AddTag tag
// @Summary AddTag
// @Produce  json
// @Tags  标签
// @Security ApiKeyAuth
// @Param body body v1.tagCreate true "新建"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /tags [post]
func AddTag(c *gin.Context) {
	var tag tagCreate
	c.BindJSON(&tag)

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	fmt.Println(tag)
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		if !models.ExistTagByName(tag.Name) {
			code = e.SUCCESS
			models.AddTag(tag.Name, tag.State, 1)
			maps := make(map[string]interface{})
			maps["name"] = tag.Name
			data["tag"] = models.GetOneTags(maps)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// EditTag byId
// @Tags  标签
// @Summary get a tag
// @Security ApiKeyAuth
// @Description get tag by ID
// @ID tagId
// @Accept  json
// @Produce  json
// @Param id path int true "tag ID"
// @param body body v1.tagCreate true "修改"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /tags/{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var tag tagCreate
	c.BindJSON(&tag)

	valid := validation.Validation{}
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			if tag.Name != "" {
				data["name"] = tag.Name
			}
			if tag.State != -1 {
				data["state"] = tag.State
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

// GetTagByID id
// @Tags  标签
// @Summary put a tag
// @Security ApiKeyAuth
// @Description put tag by ID
// @ID tagId
// @Accept  json
// @Produce  json
// @Param id path int true "tag ID"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /tags/{id} [get]
func GetTagByID(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	code := e.SUCCESS
	if models.ExistTagByID(id) {
		maps["ID"] = id
		data["tag"] = models.GetOneTags(maps)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// DelTagByID id
// @Tags  标签
// @Summary delete a tag
// @Security ApiKeyAuth
// @Description delete tag by ID
// @ID tagId
// @Accept  json
// @Produce  json
// @Param id path int true "tag ID"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /tags/{id} [delete]
func DelTagByID(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if models.ExistTagByID(id) {
		models.DelTagByID(id)
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
