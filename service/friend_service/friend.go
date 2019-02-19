package friend_service

import "Gin-Admin-Blog/models"

type Friend struct {
	Id int64

	Name  string
	Link  string
	Image string
	Desc  string
}

func (Friend) GetFriends() ([]*models.Friend, error) {

	friends, err := models.GetFriends()
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (f *Friend) ExistById() (bool, error) {
	return models.ExistFriendById(f.Id)
}

func (f *Friend) Delete() error {
	if err := models.DeleteCategory(f.Id); err != nil {
		return err
	}
	return nil
}

func (f *Friend) Add() error {
	friend := map[string]interface{}{
		"name":  f.Name,
		"desc":  f.Desc,
		"image": f.Image,
		"link":  f.Link,
	}

	if err := models.AddFriend(friend); err != nil {
		return err
	}
	return nil
}

//func (f *Friend) Update() error {
//	category := map[string]interface{}{
//		"name": c.Name,
//		"id":   c.Id,
//	}
//	if err := models.UpdateCategory(category); err != nil {
//		return err
//	}
//	return nil
//}
