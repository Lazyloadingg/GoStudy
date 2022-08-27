package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	id   int
	name string
	age  int
}

//本地数据库信息
var (
	dbUserName    = "root"
	dbPass        = "Admin123..!"
	dbHost        = "127.0.0.1"
	dbPort        = 3306
	dbCharset     = "utf8"
	dbName        = "test_go"
	userTableName = "test_user_info"
)

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("\"数据库连接成功\": %v\n", "数据库连接成功")
	}

	user := Person{
		name: "小明",
		age:  18,
	}

	result, id := addUser(user)
	fmt.Printf("result: %v\n", result)
	if result {
		fmt.Printf("\"插入id\": %v  %v\n", "插入成功", id)
		user := queryUserWithId(id)
		fmt.Printf("user: %v\n", user)
	}

	results := queryUsers(0)
	if len(results) > 0 {
		fmt.Printf("\"多行查询成功\": %v   %v\n", "多行查询成功", results)
	}

}

func initDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbUserName, dbPass, dbHost, dbPort, dbName, dbCharset)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}

func addUser(user Person) (bool, int) {
	sql := fmt.Sprintf("insert into %s(name,age) VALUES(?,?)", userTableName)
	result, err := db.Exec(sql, user.name, user.age)
	if err != nil {
		fmt.Printf("\"插入失败1\": %v\n", "插入失败")
		return false, 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("\"插入失败2\": %v\n", "插入失败")
	}
	fmt.Printf("id: %v\n", id)
	return true, int(id)
}

//单行查询
func queryUserWithId(id int) Person {
	sql := fmt.Sprintf("select name,age from %s where id= ?", userTableName)
	var user Person
	err := db.QueryRow(sql, id).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("\"查询失败%v\": %v\n", err, "查询失败")
	}
	return user
}

//多行查询
func queryUsers(id int) []Person {

	sql := fmt.Sprintf("select name,age from %s where id > ?", userTableName)
	results, err := db.Query(sql, id)
	if err != nil {
		fmt.Printf("\"查询失败\": %v\n", "查询失败")
	}
	res := make([]Person, 0)
	for results.Next() {
		var user Person
		err := results.Scan(&user.name, &user.age)
		if err != nil {
			continue
		}
		fmt.Printf("user.name: %v\n", user.name)
		res = append(res, user)
	}
	return res
}
