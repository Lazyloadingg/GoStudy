package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var (
	dbUserName    = "root"
	dbPass        = "Admin123..!"
	dbHost        = "127.0.0.1"
	dbPort        = 3306
	dbCharset     = "utf8"
	dbName        = "test_go"
	userTableName = "test_user_info"
)
var dbx *sqlx.DB

func sqlxDB() {
	err := initDBX()
	if err != nil {
		fmt.Printf("数据库启动失败err: %v\n", err)
		return
	}
	fmt.Printf("\"数据库启动成功\": %v\n", "数据库启动成功")
}
func getSqlDriver() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbUserName, dbPass, dbHost, dbPort, dbName, dbCharset)
}

func initDBX() (err error) {

	dbx, err = sqlx.Connect("mysql", getSqlDriver())
	if err != nil {
		return err
	}
	return err
}

func queryUserWithID(id int) (error, Person) {
	sql := fmt.Sprintf("select * from %s where id = ?", userTableName)

	var user Person
	err := dbx.Get(&user, sql, id)
	if err != nil {
		fmt.Printf("查询失败err: %v\n", err)
		return err, user
	} else {
		return nil, user
	}
}
