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

// ------

type AddCategoryForm struct {
	Name string `form:"name"`
}

func AddCategory(c *gin.Context) {
	appG := app.Gin{C: c}

	var form AddCategoryForm

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{Name: form.Name}

	if err := categoryService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//-------------------------------
type deleteForm struct {
	Id int64 `form:"id"`
}

func DeleteCategory(c *gin.Context) {
	appG := app.Gin{C: c}

	var form deleteForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{Id: form.Id}

	if err := categoryService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//-------------------------------
type UpdateCategoryForm struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
}

func UpdateCategory(c *gin.Context) {
	appG := app.Gin{C: c}

	var form UpdateCategoryForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{Id: form.Id, Name: form.Name}

	if err := categoryService.Update(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
