package models

import "gin_blog/dao"

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态，1删除
	CreateTime int64
}

func InsertUser(user *User) (int64, error){
	return dao.ModifyDB("insert into users(username,password,status,create_time) values(?,?,?,?)",
		user.Username,user.Password,user.Status,user.CreateTime)
}

//根据用户名查id
func QueryUserWithUsername(username string) int {
	var user User
	//row := dao.QueryRowDB("select id from users where username=?",username)
	//id := 0
	//row.Scan(&id)
	//return id
	err := dao.QueryRowDB(&user,"select id from users where username=?",username)
	if err != nil {
		return 0
	}
	return user.Id

}

func QueryUserWithParam(username, password string) int {
	//row := dao.QueryRowDB("select id from users where username=? and password=?", username, password)
	//id := 0
	//row.Scan(&id)
	//return id
	var user User
	err := dao.QueryRowDB(&user,"select id from users where username=?",username)
	if err != nil {
		return 0
	}
	return user.Id

}
