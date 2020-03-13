package main

import (
	"fmt"
	"gin_blog/dao"
	"gin_blog/routers"
)


func main() {
	err := dao.InitMySQL()
	if err != nil {
		fmt.Println("init MySQL failed, err:%v\n")
		return
	}
	r := routers.SetupRouter()
	r.Run()
}
