package models

import "github.com/jinzhu/gorm"

type Blog struct {
	Category Category `json:"category" gorm:"ForeignKey:Id"`

	//TagId int `json:"tag_id"`
	Tags []Tag `json:"tags" gorm:"many2many:blog_blog_tags;"`

	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	PageView int    `json:"page_view"`
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
