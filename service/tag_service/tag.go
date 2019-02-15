package tag_service

import "Gin-Admin-Blog/models"

type Tag struct {
	Ids []int64

	Id   int64
	Name string
}

func (Tag) GetTags() ([]*models.Tag, error) {

	tags, err := models.GetTags()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) ExistTagByIds() (bool, error) {
	return models.ExistTagByIds(t.Ids)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.Id)
}

func (t *Tag) Add() error {

	tag := map[string]interface{}{
		"name": t.Name,
	}
	if err := models.AddTag(tag); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Update() error {
	tag := map[string]interface{}{
		"name": t.Name,
		"id":   t.Id,
	}
	if err := models.UpdateTag(tag); err != nil {
		return err
	}
	return nil
}
