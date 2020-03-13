package dao

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitMySQL()(err error){
	fmt.Println("InitMySQL....")
	if db == nil {
		db, err = sql.Open("mysql","root:root123@tcp(127.0.0.1)/gin_blog")
		if err != nil {
			return
		}
	}
	return CreateTableWithUser()
}

//创建用户表
func CreateTableWithUser()(err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        create_time INT(10)
        );`

	_, err = ModifyDB(sqlStr)
	return
}

//操作数据库
func ModifyDB(sql string, args ...interface{})(int64, error){
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	count,err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

//查询
func QueryRowDB(query string, args ...interface{}) *sql.Row{
	return db.QueryRow(query, args...)
}