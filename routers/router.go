package routers

import (
	"Gin-Admin-Blog/middleware/jwt"
	"Gin-Admin-Blog/pkg/setting"
	"Gin-Admin-Blog/routers/api"
	"Gin-Admin-Blog/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gin.SetMode(setting.ServerSetting.RunMode)

	//r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("auth", api.GetAuth)

	vOne := r.Group("api/v1")
	vOne.Use(jwt.JWT())
	{

		vOne.GET("user/info", v1.GetUser)
		vOne.GET("blogs", v1.GetBlogs)

		vOne.GET("categorys", v1.GetCategorys)

		vOne.GET("tags", v1.GetTags)
	}
	return r
}
