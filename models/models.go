package models

import (
	"Gin-Admin-Blog/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
)

import _ "github.com/jinzhu/gorm/dialects/sqlite"

type Model struct {
}

var db *gorm.DB

func SetUp() {

	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.Path)

	if err != nil {
		log.Fatalf("models.SetUp err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDb() {
	defer db.Close()
}
