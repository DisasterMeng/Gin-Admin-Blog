package models

import (
	"Gin-Admin-Blog/pkg/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	ID          int64  `gorm:"primary_key" json:"id"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	IsSuperuser bool   `json:"is_superuser"`
	IsActive    bool   `json:"is_active"`
	Email       string `json:"email"`
}

func (User) TableName() string {
	return "user_user"
}

func CheckAuth(username, password string) (bool, error) {
	var user User
	e := db.Where(User{Username: username}).First(&user).Error

	if e != nil && e != gorm.ErrRecordNotFound {
		return false, e
	}

	if !user.IsSuperuser || !user.IsActive {
		return false, errors.New("user is't active or superuser")
	}
	return utils.CheckPassword(password, user.Password)

}

func GetUser(username string) User {
	var user User
	db.Where(User{Username: username}).First(&user)
	return user
}

func GetUserByID(id int64) User {
	var user User
	db.First(&user, id)
	return user
}
