package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func initDB() (err error) {
	db, err = sql.Open("", "")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}
