package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/pkg/upload"
	"Gin-Admin-Blog/service/friend_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFriends(c *gin.Context) {
	appG := app.Gin{C: c}

	friendService := friend_service.Friend{}

	friends, error := friendService.GetFriends()

	if error != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"friends": friends,
	})

}

// ------

type AddFriendForm struct {
	Name string `form:"name" valid:"Required"`
	Link string `form:"link" valid:"Required"`
	Desc string `form:"desc"`
}

func AddFriend(c *gin.Context) {
	appG := app.Gin{C: c}

	var form AddFriendForm

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	file, image, err := c.Request.FormFile("image")

	saveUrl := ""

	//如果存在图片
	if image != nil {
		up := upload.NewUpload(upload.TYPE_FRIEND)

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

	friendService := friend_service.Friend{Name: form.Name, Desc: form.Desc, Image: saveUrl, Link: form.Link}

	if err = friendService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FRIEND_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

///-------------

func DeleteFriend(c *gin.Context) {
	appG := app.Gin{C: c}

	var form IdForm

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	friendService := friend_service.Friend{Id: form.Id}

	if err := friendService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FRIEND_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
