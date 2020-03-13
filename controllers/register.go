package controllers

import (
	"fmt"
	"gin_blog/gb"
	"gin_blog/models"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RegisterGet(c *gin.Context){
	c.HTML(http.StatusOK,"register.html",gin.H{"title":"注册页"})

}

func RegisterPost(c *gin.Context){
	//取出请求数据
	//数据校验 有效性
	//判断注册是否重复 拿到数据库对比
	//写入数据库
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:",id)
	if id > 0 {
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"用户名已经存在"})
		return
	}
	password = utils.MD5(password)
	gb.Logger.Debug(zap.String("md5",password))

	user := models.User{
		Username:username,
		Password:password,
		Status:0,
		CreateTime:time.Now().Unix(),
	}
	_,err := models.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"code":0,"message":""})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":1,"message":"注册成功"})
	}
}