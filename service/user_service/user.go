package user_service

import "Gin-Admin-Blog/models"

type User struct {
	ID int64
}

func (u *User) Get() models.User {
	return models.GetUserByID(u.ID)
}
