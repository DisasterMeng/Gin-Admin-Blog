package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/service/blog_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBlogs(c *gin.Context) {
	appG := app.Gin{C: c}

	blogService := blog_service.Blog{PageSize: 1, PageNum: 0}
	blogs, _ := blogService.GetAll()

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"blogs": blogs,
	})

}
