package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/pkg/upload"
	"Gin-Admin-Blog/service/blog_service"
	"Gin-Admin-Blog/service/category_service"
	"Gin-Admin-Blog/service/tag_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogListForm struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}

func GetBlogs(c *gin.Context) {
	appG := app.Gin{C: c}

	var form BlogListForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	blogService := blog_service.Blog{PageSize: form.PageSize, PageNum: form.PageNum}
	blogs, _ := blogService.GetAll()

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"blogs": blogs,
	})

}

//-----------------------
type AddBlogForm struct {
	CategoryId int64   `form:"category_id" valid:"Required;Min(1)"`
	Title      string  `form:"title" valid:"Required;MaxSize(100)"`
	Content    string  `form:"content" valid:"Required;MaxSize(65535)"`
	TagIds     []int64 `form:"tag_ids"`
	//SummaryImg string `form:"summary_img" valid:"Required"`
}

func AddBlog(c *gin.Context) {
	appG := app.Gin{C: c}
	var form AddBlogForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := category_service.Category{Id: form.CategoryId}
	exist, err := categoryService.ExistById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_CATEGORY_FAIL, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	tagService := tag_service.Tag{Ids: form.TagIds}
	exist, err = tagService.ExistTagByIds()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_TAG_FAIL, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	file, image, err := c.Request.FormFile("summary_img")

	saveUrl := ""

	//如果存在图片
	if image != nil {
		up := upload.NewUpload(upload.TYPE_SUMMARY)

		imageName := up.GetImageName(image.Filename)

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
			return
		}

		err = up.CheckImage()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
			return
		}

		if err := up.SaveUploadedFile(image); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
			return
		}
		saveUrl = up.GetImageFullPath()
	}

	blogService := blog_service.Blog{TagIds: form.TagIds, CategoryId: form.CategoryId, Content: form.Content, Title: form.Title, SummaryImg: saveUrl}

	if err = blogService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_BLOG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//-------------------------------
type DeleteForm struct {
	Id int64 `form:"id"`
}

func DeleteBlog(c *gin.Context) {
	appG := app.Gin{C: c}

	var form DeleteForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	blogService := blog_service.Blog{Id: form.Id}

	if err := blogService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_BLOG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
