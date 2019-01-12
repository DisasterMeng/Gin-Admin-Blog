package category_service

import "Gin-Admin-Blog/models"

type Category struct {
}

func (Category) GetCategorys() ([]*models.Category, error) {

	categories, err := models.GetCategorys()
	if err != nil {
		return nil, err
	}

	return categories, nil

}
