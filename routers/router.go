package routers

import (
	"gin_blog/controllers"
	"github.com/gin-gonic/gin"
)

//func
func SetupRouter() *gin.Engine{
	r := gin.Default()
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