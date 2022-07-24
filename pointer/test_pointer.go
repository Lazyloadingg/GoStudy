package main

import "fmt"

func main() {

	var a = 10 //a为值类型变量
	var b = &a //b为指针变量
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	//new函数分配内存并返回一个指针
	c := new(int)
	fmt.Printf("c: %v\n", c)       //直接打印是地址，因为`new`返回的指针，和上边的b类似
	fmt.Printf("(*c): %v\n", (*c)) //要想取值就要用`*`
	*c = 10
	fmt.Printf("修改后(*c): %v\n", (*c)) //要想取值就要用`*`

}
