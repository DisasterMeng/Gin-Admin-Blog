package routers

import (
	"Gin-Admin-Blog/middleware/jwt"
	"Gin-Admin-Blog/pkg/setting"
	"Gin-Admin-Blog/routers/api"
	"Gin-Admin-Blog/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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
	r.StaticFS(setting.AppSetting.UploadDir, http.Dir(setting.AppSetting.UploadDir))
	r.POST("auth", api.GetAuth)
	r.POST("upload-summary", v1.UploadSummary)

	vOne := r.Group("api/v1")
	vOne.Use(jwt.JWT())
	{

		vOne.GET("user/info", v1.GetUser)

		vOne.GET("blog/list", v1.GetBlogs)
		vOne.POST("blog/add", v1.AddBlog)
		vOne.POST("blog/delete", v1.DeleteBlog)

		vOne.GET("category/list", v1.GetCategorys)
		vOne.POST("category/delete", v1.DeleteCategory)
		vOne.POST("category/add", v1.AddCategory)
		vOne.POST("category/update", v1.UpdateCategory)

		vOne.GET("tag/list", v1.GetTags)
		vOne.POST("tag/delete", v1.DeleteTag)
		vOne.POST("tag/add", v1.AddTag)
		vOne.POST("tag/update", v1.UpdateTag)

		r.POST("md2html", v1.MdToHtml)

	}
	return r
}
