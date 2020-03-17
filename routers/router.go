package routers

import (
	"gin_blog/controllers"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

//func
func SetupRouter() *gin.Engine{
	r := gin.New()

	r.Use(logger.GinLogerr(logger.Logger),logger.GinRecovery(logger.Logger, true))

	r.SetFuncMap(template.FuncMap{
		"timeStr":func(timestamp int64) string{
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})
	r.Static("/static","static")
	r.LoadHTMLGlob("views/*")

	//注册
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register", controllers.RegisterPost)

	r.GET("/login", controllers.LoginGet)
	r.POST("/login",controllers.LoginPost)

	r.GET("/",controllers.HomeGet)

	article := r.Group("/article")
	{
		article.GET("/add",controllers.AddArticleGet)
		article.POST("/add",controllers.AddArticlePost)
	}
	return r
}