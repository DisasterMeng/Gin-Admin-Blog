package upload

import (
	"Gin-Admin-Blog/pkg/file"
	"Gin-Admin-Blog/pkg/setting"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadLocal struct {
	Type int
	Name string
}

func (u *UploadLocal) GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	// 这里可以对文件名字进行处理
	u.Name = fileName + ext
	return u.Name
}

func (u *UploadLocal) GetImagePath() string {

	switch u.Type {
	case TYPE_USER:
		return setting.AppSetting.UploadUserDir
	case TYPE_FRIEND:
		return setting.AppSetting.UploadFriendDir
	case TYPE_SUMMARY:
		return setting.AppSetting.UploadSummaryDir
	}

	return ""
}

func (u *UploadLocal) CheckImage() error {
	var src = u.GetImagePath()
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

func (u *UploadLocal) SaveUploadedFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(u.GetImageFullPath())
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, src)
	return nil
}

func (u *UploadLocal) GetImageFullPath() string {
	return u.GetImagePath() + u.Name
}

func (u *UploadLocal) GetImageFullUrl() string {
	return setting.AppSetting.PrefixUrl + "/" + u.GetImagePath() + u.Name
}
