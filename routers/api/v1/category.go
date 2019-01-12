package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/service/category_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategorys(c *gin.Context) {
	appG := app.Gin{C: c}

	categoryService := category_service.Category{}

	categories, error := categoryService.GetCategorys()

	if error != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"categories": categories,
	})

}
