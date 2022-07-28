package main

import (
	"fmt"
	"sync"
)

//定义全局的等待组
var g sync.WaitGroup

func main() {

	// hello()
	// fmt.Printf("\"你好世界\": %v\n", "你好世界")

	/*
		打印顺序
		1. 自律使你自由
		2. 你好世界
	*/
	// ---------------分割线----------------

	//开启一个`goroutine`执行函数
	// go hello()
	// fmt.Printf("\"你好世界\": %v\n", "你好世界")
	// time.Sleep(time.Duration(3) * time.Second)
	// fmt.Printf("\"醒了\": %v\n", "醒了")
	/*
		打印顺序
		1. 你好世界
		2. 自律使你自由

		解释：
		1. 因为开启`goroutine`需要一定的开销，而此时main的goroutine还在继续向下执行，因此会先执行`你好世界`
		2. 之所以等待3秒时候因为main goroutine执行完时，hello()的go goroutine还没创建成功，因此程序结束被销毁，hello还没执行
		3. 这种睡眠等待方式是不优雅的
	*/

	//使用WaitGroup
	// g.Add(1)//等级一个goroutine
	// go hello()
	// fmt.Printf("\"你好世界\": %v\n", "你好世界")
	// g.Wait()//

	//写一个iOS常写的GCD任务顺序的go实现
	for i := 0; i < 10; i++ {
		g.Add(1)
		go showNum(i)

	}
	g.Wait()

}

func hello() {
	fmt.Printf("\"自律使你自由\": %v\n", "自律使你自由")
	//使用等待组时才调用
	g.Done() //告诉等待组要等待的任务完成，类似iOS中信号量
}

func showNum(i int) {
	defer g.Done()
	fmt.Printf("看看i: %v\n", i)

}
