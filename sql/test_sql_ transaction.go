package main

import (
	"fmt"
)

//事务操作，保证一组操作同时完成或同时失败，没有中间状态
func transactionExce() {

	tx, errTx := db.Begin() //开启事务

	if errTx != nil {
		fmt.Printf("\"事务开启失败\": %v\n", "事务开启失败")
		tx.Rollback()
		return
	}

	//演示事务直接拼接sql不写那么复杂了>_<!

	sql := fmt.Sprintf("update %v set name = ? where id = ?", userTableName)

	ret, err := tx.Exec(sql, "小刚1", 2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		tx.Rollback()
		return
	}

	row, errAff := ret.RowsAffected()
	if errAff != nil {
		fmt.Printf("errAff: %v\n", errAff)
		tx.Rollback()
		return
	}

	fmt.Printf("更新姓名成功row: %v\n", row)
	sql1 := fmt.Sprintf("update %v set age = ? where id = ?", userTableName)

	ret1, err1 := tx.Exec(sql1, 21, 2)

	if err1 != nil {
		fmt.Printf("err1: %v\n", err1)
		tx.Rollback()
		return
	}

	row1, errAff1 := ret1.RowsAffected()
	if errAff1 != nil {
		fmt.Printf("errAff1: %v\n", errAff1)
		tx.Rollback()
		return
	}
	fmt.Printf("更新年龄成功row1: %v\n", row1)

	if row == 1 && row1 == 1 {
		err2 := tx.Commit()
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		} else {
			fmt.Printf("\"事务操作成功\": %v\n", "事务操作成功")
		}
	} else {
		fmt.Printf("\"事务操作失败\": %v\n", "事务操作失败")
	}

}
