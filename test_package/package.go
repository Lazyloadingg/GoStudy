package main

import (
	"fmt"
	"runtime"
)

//包级别变量，会在`init`函数之前初始化完成
var PI = 3.1415
var b int

const a = 10

func main() {

	// init() 手动调用会报错
	fmt.Printf("\"main\": %v\n", "main")
	Say()
}

/*
1. 每个go文件中都可以定义多个`init`函数，这个函数不能被手动调用，会被系统调用
2. init调用顺序会按照引入顺序来执行，所有引入包init执行完后会执行`main`的`init`
3. init在main函数之前执行
*/

func init() {
	//此时打印这些变量会发现已经有值了
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("PI: %v\n", PI)

	fmt.Printf("初始函数1\n")
}
func init() {
	fmt.Printf("初始函数2\n")
	cus()
}

func cus() {
	pc, file, line, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("file: %v\n", file)
	fmt.Printf("line: %v\n", line)
	fmt.Printf("\"被调用了\": %v\n", "被调用了")
}
