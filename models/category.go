package models

import "github.com/jinzhu/gorm"

type Category struct {
	Model

	//Id   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (Category) TableName() string {
	return "blog_category"
}

func ExistTagById(id int64) (bool, error) {
	var category Category
	err := db.First(&category, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}
	if category.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetCategorys() ([]*Category, error) {
	var categorys []*Category
	err := db.Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categorys, nil
}

func GetCategory(id int64) (*Category, error) {
	var category Category
	err := db.First(&category, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &category, nil
}

func DeleteCategory(id int64) error {
	model := Model{Id: id}
	category := Category{Model: model}

	if err := db.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func AddCategory(data map[string]interface{}) error {
	category := Category{
		Name: data["name"].(string),
	}
	if err := db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategory(data map[string]interface{}) error {
	model := Model{Id: data["id"].(int64)}
	category := Category{
		Model: model,
	}
	name := data["name"].(string)
	if err := db.Model(&category).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}
