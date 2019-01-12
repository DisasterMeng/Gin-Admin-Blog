package models

import "github.com/jinzhu/gorm"

type Category struct {
	Id   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (Category) TableName() string {
	return "blog_category"
}

func GetCategorys() ([]*Category, error) {
	var categorys []*Category
	err := db.Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categorys, nil
}
