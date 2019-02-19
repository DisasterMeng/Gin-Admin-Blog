package models

import "github.com/jinzhu/gorm"

type Friend struct {
	Model

	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

func (Friend) TableName() string {
	return "friend_friend"
}

func GetFriends() ([]*Friend, error) {
	var friends []*Friend
	err := db.Find(&friends).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return friends, nil
}

func ExistFriendByIds(ids []int64) (bool, error) {
	var friends []*Friend
	err := db.Where(ids).Find(&friends).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func ExistFriendById(id int64) (bool, error) {
	var friend Friend
	err := db.First(&friend, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}
	if friend.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetFriendsByIds(ids []int64) ([]Friend, error) {
	var friends []Friend
	err := db.Where(ids).Find(&friends).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return friends, nil
}

func DeleteFriend(id int64) error {
	model := Model{Id: id}
	tag := Tag{Model: model}

	if err := db.Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}

func AddFriend(data map[string]interface{}) error {
	friend := Friend{
		Name: data["name"].(string),
	}
	if err := db.Create(&friend).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFriend(data map[string]interface{}) error {
	model := Model{Id: data["id"].(int64)}
	tag := Tag{
		Model: model,
	}
	name := data["name"].(string)
	if err := db.Model(&tag).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}
