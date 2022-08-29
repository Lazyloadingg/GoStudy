package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Personx struct {
	ID   int
	Name string
	Age  int
}

var dbx *sqlx.DB

func initDBX() (err error) {
	dbx, err = sqlx.Connect("mysql", getSqlDriver())
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return err
}

func sqlxTest() {

	// https://www.liwenzhou.com/posts/Go/sqlx/
	// sqlx是一个系统标准库`sql`的超集，提供了很多方便的扩展
	err := initDBX()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	user := sqlxQuery(3)
	fmt.Printf("sqlx查询成功: %v\n", user)

	users := sqlxQueryUsers(10)
	fmt.Printf("sqlx多行查询成功users: %v\n", users)
}

//单行查询
func sqlxQuery(id int) Personx {
	//因为上个文件中定义的`Person`字段为小写，不能跨包在sqlx中访问，因此重新定义一个Personx
	var user Personx
	sql := fmt.Sprintf("select id , name ,age from %s where id = ?", userTableName)
	//相比sql，sqlx查询结果可以直接传入一个结果目标类型并将查询结果赋值进去，而sql还需要手动一个一个字段赋值
	err := dbx.Get(&user, sql, id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("\"sqlx单行查询失败\": %v\n", "sqlx单行查询失败")
	}
	return user

}

//sqlx多行查询
func sqlxQueryUsers(id int) []Personx {
	dbx.Stats()
	sql := fmt.Sprintf("select id,name,age from %s where id > ?", userTableName)
	var users []Personx
	//相比sql  可以直接传一个切片进去返回目标结果集，而不用自己手动创建切片，通过Next遍历添加
	err := dbx.Select(&users, sql, id)
	if err != nil {
		fmt.Printf("\"sqlx多行查询失败\": %v\n", "sqlx多行查询失败")
		return nil
	}
	return users
}
