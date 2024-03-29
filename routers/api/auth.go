package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gsgo/models"
	"gsgo/pkg/e"
	"gsgo/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Tags auth
// @Produce  json
// @Param body body api.auth true "登录"
// @Success 200 {string} string ""
// @Failure 500 {string} string ""
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	var user auth
	c.BindJSON(&user)
	username := user.Username
	password := user.Password
	// fmt.Println(username, password)

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		id := models.CheckAuth(username, password)
		if id != 0 {
			token, err := util.GenerateToken(username, id)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
