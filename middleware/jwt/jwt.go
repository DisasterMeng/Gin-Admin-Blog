package jwt

import (
	"Gin-Admin-Blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("token")

		if token != "TembinCrawlerAuth" { //先硬编码
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data":    data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
