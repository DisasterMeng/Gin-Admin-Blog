package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	Model

	CategoryId int64    `json:"category_id"`
	Category   Category `json:"category" gorm:"ForeignKey:CategoryId"`

	Tags []Tag `json:"tags" gorm:"many2many:blog_blog_tags;"`

	Title      string `json:"title"`
	Content    string `json:"content"`
	PageView   int    `json:"page_view"`
	SummaryImg string `json:"summary_img"`
}

func (Blog) TableName() string {
	return "blog_blog"
}

func GetBlogs(pageNum int, pageSize int) ([]*Blog, error) {
	var blogs []*Blog
	err := db.Preload("Category").Preload("Tags").Offset(pageNum).Limit(pageSize).Find(&blogs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return blogs, nil
}

func AddBlog(data map[string]interface{}) error {

	//var tags []Tag
	////var category Category
	//
	//if tagIds, ok := data["tag_ids"].([]int64); ok {
	//	db.Where(tagIds).Find(&tags)
	//}

	tags, _ := GetTagByIds(data["tag_ids"].([]int64))
	blog := Blog{
		CategoryId: data["category_id"].(int64),
		Title:      data["title"].(string),
		Content:    data["content"].(string),
		Tags:       tags,
		SummaryImg: data["summary_img"].(string),
	}
	if err := db.Create(&blog).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBlog(x int64) error {
	model := Model{Id: x}
	blog := Blog{Model: model}

	if err := db.Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}
