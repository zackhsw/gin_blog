package controllers

import (
	"fmt"
	"gin_blog/gb"
	"gin_blog/models"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func LoginGet(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{"tile":"登录页"})
}

func LoginPost(c *gin.Context){
	//取出请求数据
	//校验用户名密码
	//返回响应
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:",username,"password",password)
	gb.Logger.Debug("username",username,"password",password)

	id := models.QueryUserWithParam(username,utils.MD5(password))
	fmt.Println("id:",id)
	if id > 0 {
		//给响应添加cookie
		session := sessions.Default(c)
		session.Set("login_user",username)
		session.Save()
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"登录成功"})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"登录失败"})
	}
}