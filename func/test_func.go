package main

import (
	"fmt"
)

//定义函数类型，写法也和swift类似
type test func(int, int) int

func main() {
	// fmt.Printf("%v", test_func(2, 3))
	// fmt.Printf("%v", test_func1())
	// fmt.Printf("%v", test_sum(1, 2, 3, 4, 5))

	//函数类型变量
	// var fn test = test_func
	// fmt.Printf("%v", fn(1, 2))

	// test_1()
	// test_closure()
	// test_defer()
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())

	// x := 1
	// y := 2
	// defer calc("AA", x, calc("A", x, y))
	// x = 10
	// defer calc("BB", x, calc("B", x, y))
	// y = 20

	test_error()
	fmt.Printf("\"处理完异常\": %v\n", "处理完异常")

}

//有参数有返回值
func test_func(a int, b int) (ret int) {
	ret = a + b + 1
	return ret
}

//无参数有返回值
func test_func1() int {
	return 10
}

//可变参数，args本质是一个切片；这里可变参数和swift里类似
func test_sum(args ...int) int {
	fmt.Println(args)
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum

}

//匿名函数，整体也和swift很像，上手很容易
func test_1() {

	//匿名函数因为没有函数名无法直接调用，因此需要保存到一个变量中进行调用
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Printf("%v\n", add(1, 2))

	//或者立即执行，这的res接收的不是函数变量，而是函数返回值，和上边是不一样的
	res := func(a int, b int) int {
		return a + b
	}(3, 4)
	fmt.Printf("res: %v\n", res)
}

//闭包，和swift也类似，都是封装了函数+函数调用环境，他是有状态的函数（包含内部使用的变量等）
func test_closure() {
	//fn为闭包
	var fn = test_closure1()
	fmt.Printf("%v", fn(10)) //10
	fmt.Printf("%v", fn(20)) //30
	fmt.Printf("%v", fn(30)) //60
	/**
	因为fn是个变量已经接收了test_closure1函数返回的函数，内部返回的函数捕获了局部变量`res`，
	因此在变量fn生命周期内`res`都会一直在内存中，因此多次调用结果是递增的
	*/
}

func test_closure1() func(int) int {
	var res int
	return func(y int) int {
		res += y
		return res
	}
}

/**
defer 被defer修饰的语句会在作用域结束前最后执行，如果有多个被defer修饰的语句，先defer的语句，后执行，但是所有defer语句都晚于其他语句执行
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：
*/
func test_defer() {
	fmt.Printf("\"开始\": %v\n", "开始")
	defer fmt.Printf("\"1\": %v\n", "1") //先被defer，最后执行
	defer fmt.Printf("\"2\": %v\n", "2") //
	defer fmt.Printf("\"3\": %v\n", "3") //defer都在其他语句`开始``结束`之后
	fmt.Printf("\"结束\": %v\n", "结束")
	/** 打印结果
	"开始": 开始
	"结束": 结束
	"3": 3
	"2": 2
	"1": 1
	*/
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)

	return ret
}

func test_error() {
	fmt.Printf("\"进入函数\": %v\n", "进入函数")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("捕获异常 %v", err)
		}
	}()
	panic("异常了")
	fmt.Printf("\"退出函数\": %v\n", "退出函数")
}
