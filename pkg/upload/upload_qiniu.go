package upload

import (
	"Gin-Admin-Blog/pkg/file"
	"Gin-Admin-Blog/pkg/setting"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadQiNiu struct {
	Type int
	Name string
}

func (u *UploadQiNiu) GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	// 这里可以对文件名字进行处理
	u.Name = fileName + ext
	return u.Name
}

func (u *UploadQiNiu) GetImagePath() string {

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

func (u *UploadQiNiu) CheckImage() error {
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

func (u *UploadQiNiu) SaveUploadedFile(file *multipart.FileHeader) error {

	putPolicy := storage.PutPolicy{
		Scope: setting.QinNiuSetting.QinNiuBucketName,
	}

	mac := qbox.NewMac(setting.QinNiuSetting.QinNiuAccessKey, setting.QinNiuSetting.QinNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": u.Name,
		},
	}
	open, e := file.Open()
	if e != nil {
		return e
	}
	defer open.Close()

	err := formUploader.Put(context.Background(), &ret, upToken, u.GetImageFullPath(), open, file.Size, &putExtra)
	if err != nil {
		return err
	}
	return nil
}

func (u *UploadQiNiu) GetImageFullPath() string {
	return u.GetImagePath() + u.Name
}

func (u *UploadQiNiu) GetImageFullUrl() string {
	return "http://" + setting.QinNiuSetting.QinBucketDomain + "/" + u.GetImageFullPath()
}
