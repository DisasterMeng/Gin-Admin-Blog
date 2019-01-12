package tag_service

import "Gin-Admin-Blog/models"

type Tag struct {
}

func (Tag) GetTags() ([]*models.Tag, error) {

	tags, err := models.GetTags()
	if err != nil {
		return nil, err
	}

	return tags, nil
}
