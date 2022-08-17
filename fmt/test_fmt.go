package main

import "fmt"

func main() {
	//不同于iOS ，go中打印有个标准库，提供了很多种打印方法，

	fmt.Print("直接将内容输出不附加任何内容")
	fmt.Print("再打印一条会发现和上一句连着在同一行")
	fmt.Printf("可以格式化内容比如：占位符%v", "我是占位符")
	fmt.Println("会自动添加换行")
	fmt.Println("和上边不在同一行了")
}
