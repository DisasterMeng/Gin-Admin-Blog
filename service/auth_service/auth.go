package auth_service

import "Gin-Admin-Blog/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) Get() models.User {
	return models.GetUser(a.Username)
}
