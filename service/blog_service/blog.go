package blog_service

import "Gin-Admin-Blog/models"

type Blog struct {
	PageSize int
	PageNum  int

	CategoryId int64

	TagIds []int64

	Title      string
	Content    string
	PageView   int
	SummaryImg string
	Id         int64
}

func (b *Blog) GetAll() ([]*models.Blog, error) {

	blogs, err := models.GetBlogs(b.PageNum, b.PageSize)

	if err != nil {
		return nil, err
	}

	return blogs, nil

}

func (b *Blog) Delete() error {

	if err := models.DeleteBlog(b.Id); err != nil {
		return err
	}

	return nil
}

func (b *Blog) Add() error {
	blog := map[string]interface{}{
		"tag_ids":     b.TagIds,
		"title":       b.Title,
		"summary_img": b.SummaryImg,
		"content":     b.Content,
		"category_id": b.CategoryId,
	}

	if err := models.AddBlog(blog); err != nil {
		return err
	}

	return nil
}
