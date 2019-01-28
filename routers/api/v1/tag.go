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

//-------------------------------
type DeleteTagForm struct {
	Id int64 `form:"id"`
}

func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}

	var form DeleteTagForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{Id: form.Id}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// ------

type AddTagForm struct {
	Name string `form:"name"`
}

func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}

	var form AddTagForm

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{Name: form.Name}

	if err := tagService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
