package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"github.com/gin-gonic/gin"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
)

type MarkdownForm struct {
	Content string `form:"content"`
}

func MdToHtml(c *gin.Context) {
	appG := app.Gin{C: c}

	var form MarkdownForm

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	bytes := blackfriday.Run([]byte(form.Content))

	appG.Response(http.StatusOK, e.SUCCESS, string(bytes))

}
