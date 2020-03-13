package controllers

import (
	"gin_blog/gb"
	"gin_blog/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	gb.Logger.Debug("loginUser",loginUser)
	if loginUser != nil {
		return true
	}else{
		return false
	}
}

func HomeGet(c *gin.Context){
	isLogin := GetSession(c)
	session := sessions.Default(c)
	username := session.Get("login_user").(string)
	page := 1
	articleList, err := models.QueryCurrUserArticleWithPage(username, page)
	if err != nil {
		gb.Logger.Error("models.QueryCurrUserArticleWithPage failed","error:",err)
	}
	gb.Logger.Debug("models.QueryCurrUserArticleWithPage,",articleList)
	data := models.GenHomeBlocks(articleList,isLogin)
	gb.Logger.Debug("models.GenHomeBlocks","data",data)
	c.HTML(http.StatusOK,"home.html",gin.H{"IsLogin":isLogin,"Data":data})
}

