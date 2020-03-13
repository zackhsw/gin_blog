package controllers

import (
	"gin_blog/gb"
	"gin_blog/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddArticleGet(c *gin.Context) {
	isLogin := GetSession(c)
	c.HTML(http.StatusOK,"write_article.html",gin.H{"IsLogin":isLogin})
}

func AddArticlePost(c *gin.Context) {
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	session := sessions.Default(c)
	currUserValue := session.Get("login_user")
	currentUser := currUserValue.(string)

	gb.Logger.Debug("title",title,"tags",tags)

	art := models.Article{
		Title:title,
		Tags:tags,
		Short:short,
		Content:content,
		Author:currentUser,
		CreateTime:time.Now().Unix(),
	}
	_, err := models.AddArticle(art)

	response := gin.H{}
	if err == nil {
		response = gin.H{"code":1,"message": "ok"}
	}else {
		response = gin.H{"code":0, "message":"error"}
	}
	c.JSON(http.StatusOK,response)
}