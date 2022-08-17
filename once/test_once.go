package main

import (
	"fmt"
	"sync"
)

type Person struct{}

//定义一个变量
var instance *Person
var once sync.Once

func main() {
	fmt.Printf("shared(): %p\n", shared())
	fmt.Printf("shared1(): %p\n", shared())
	fmt.Printf("shared2(): %p\n", shared())

	//最终可以看到打印结果三次获取的对象地址是一样的，并且`shared`中，实例化代码里的`执行`语句只打印了一次
}

//获取单例对象
func shared() *Person {
	//go中通过sync中的once保证代码只执行一次，once是一个结构体，内部只有一个go函数
	once.Do(func() {
		instance = &Person{}
		fmt.Printf("\"执行\": %v\n", "执行")
	})

	return instance

}
