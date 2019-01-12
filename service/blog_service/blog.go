package blog_service

import "Gin-Admin-Blog/models"

type Blog struct {
	PageSize int
	PageNum  int
}

func (b *Blog) GetAll() ([]*models.Blog, error) {

	blogs, err := models.GetBlogs(b.PageNum, b.PageSize)

	if err != nil {
		return nil, err
	}

	return blogs, nil

}
