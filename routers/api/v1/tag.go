package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/service/tag_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}

	tagService := tag_service.Tag{}

	tags, error := tagService.GetTags()

	if error != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"tags": tags,
	})

}
