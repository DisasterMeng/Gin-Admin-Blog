package upload

import (
	"Gin-Admin-Blog/pkg/file"
	"Gin-Admin-Blog/pkg/setting"
	"log"
	"mime/multipart"
	"strings"
)

type UploadImage interface {
	GetImageName(string) string
	GetImagePath() string
	CheckImage() error
	SaveUploadedFile(*multipart.FileHeader) error
	GetImageFullUrl() string
	GetImageFullPath() string
}

const (
	TYPE_USER = iota
	TYPE_FRIEND
	TYPE_SUMMARY
)

func NewUpload(t int) UploadImage {
	if setting.ServerSetting.RunMode == "debug" {
		return &UploadLocal{Type: t}
	} else {
		return &UploadQiNiu{Type: t}
	}
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	return size/1024/1024 <= setting.AppSetting.ImageMaxSize
}
