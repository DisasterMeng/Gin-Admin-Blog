package models

import (
	"Gin-Admin-Blog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

import _ "github.com/jinzhu/gorm/dialects/sqlite"

type JsonTime time.Time

//实现它的json序列化方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Model struct {
	Id       int64     `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

var db *gorm.DB

func SetUp() {

	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.Path)

	if err != nil {
		log.Fatalf("models.SetUp err: %v", err)
	}

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

// updateTimeStampForCreateCallback will set `Created`, `Modified` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("Created"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("Modified"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `Modified` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("Modified", time.Now())
	}
}

func CloseDb() {
	defer db.Close()
}
