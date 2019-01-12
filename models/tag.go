package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Id   int64  `json:"id" gorm:"primary_key:true"`
	Name string `json:"name"`
}

func (Tag) TableName() string {
	return "blog_tag"
}

func GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := db.Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}
