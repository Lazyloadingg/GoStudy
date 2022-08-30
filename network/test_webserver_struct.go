package main

type Person struct {
	Name string
	Age  int
	ID   int `db:"id"`
}
