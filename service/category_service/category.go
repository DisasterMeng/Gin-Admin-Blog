package category_service

import "Gin-Admin-Blog/models"

type Category struct {
	Id   int64
	Name string
}

func (Category) GetCategorys() ([]*models.Category, error) {

	categories, err := models.GetCategorys()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *Category) ExistById() (bool, error) {
	return models.ExistTagById(c.Id)
}

func (c *Category) Delete() error {
	if err := models.DeleteCategory(c.Id); err != nil {
		return err
	}
	return nil
}

func (c *Category) Add() error {
	category := map[string]interface{}{
		"name": c.Name,
	}

	if err := models.AddCategory(category); err != nil {
		return err
	}
	return nil
}

func (c *Category) Update() error {
	category := map[string]interface{}{
		"name": c.Name,
		"id":   c.Id,
	}
	if err := models.UpdateCategory(category); err != nil {
		return err
	}
	return nil
}
