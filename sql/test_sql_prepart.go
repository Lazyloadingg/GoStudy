package main

import "fmt"

//预处理示例
func prepareExam() {
	fmt.Printf("\"预处理操作\": %v\n", "预处理操作")
	fmt.Printf("db: %v\n", db)

	//一个sql语句分为命令部分和数据部分，普通sql语句会用占位符替换的方式拼接好完整的sql语句
	//然后交给数据库执行

	//预处理的方式会先将命令部分发给数据库进行编译，进行预处理并返回一个状态

	res := prepareQuery(0)
	if len(res) > 0 {
		fmt.Printf("\"预处理查询成功\": %v\n%v\n", "查询成功", res)
	}
}

func prepareQuery(id int) []Person {
	//
	sql := fmt.Sprintf("select id,name ,age from %s where id > ?", userTableName)
	st, err := db.Prepare(sql)
	if err != nil {
		fmt.Printf("预处理查询失败err: %v\n", err)
		return nil
	}
	defer st.Close()

	res := make([]Person, 0)
	rows, err := st.Query(id)
	if err != nil {
		fmt.Printf("预处理查询失败err: %v\n", err)
		return res
	}

	for rows.Next() {
		var user Person
		rows.Scan(&user.id, &user.name, &user.age)
		res = append(res, user)
	}
	return res

}
