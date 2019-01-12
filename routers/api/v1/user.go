package v1

import (
	"Gin-Admin-Blog/pkg/app"
	"Gin-Admin-Blog/pkg/e"
	"Gin-Admin-Blog/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.GetInt64("userID")
	user := user_service.User{ID: id}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"user": user.Get(),
	})

}
