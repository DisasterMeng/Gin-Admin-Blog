package routers

import (
	"Gin-Admin-Blog/middleware/jwt"
	"Gin-Admin-Blog/pkg/setting"
	"Gin-Admin-Blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vOne := r.Group("api/v1")
	vOne.GET("test", v1.AAAA)
	vOne.Use(jwt.JWT())
	{
		//vOne.GET("test",v1.AAAA)
	}
	return r
}
