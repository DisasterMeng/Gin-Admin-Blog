package api

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/pkg/utils"
	"Gin-Admin-Blog/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	user := authService.Get()
	token, err := utils.GenerateToken(user.ID, user.Username, user.Email)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.C.SetCookie("BLOG", token, 60*60*24, "/", "127.0.0.1:5050", false, false)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
