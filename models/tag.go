package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	//Id   int64  `json:"id" gorm:"primary_key:true"`
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

func ExistTagByIds(ids []int64) (bool, error) {
	var tags []*Tag
	err := db.Where(ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func GetTagByIds(ids []int64) ([]Tag, error) {
	var tags []Tag
	err := db.Where(ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func DeleteTag(id int64) error {
	model := Model{Id: id}
	tag := Tag{Model: model}

	if err := db.Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}

func AddTag(data map[string]interface{}) error {
	tag := Tag{
		Name: data["name"].(string),
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTag(data map[string]interface{}) error {
	model := Model{Id: data["id"].(int64)}
	tag := Tag{
		Model: model,
	}
	name := data["name"].(string)
	if err := db.Model(&tag).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}
