package main

import "fmt"

func main() {

	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()
}

func test1() {
	//定义一个通道
	var cn chan int
	fmt.Printf("cn: %v\n", cn)

	//初始化通道，缓冲区为1
	cn = make(chan int, 1)
	fmt.Printf("cn: %v\n", cn)

	//通道接收值，或者说向通道发送值，
	cn <- 10
	//继续接收，但是这样写会报错`deadlock`死锁，因为通道缓冲区为1，通道接收一个值后，在通道未将这个值发送出去之前，无法再次接收，
	//简单讲就是通道无法存储超过缓冲区的值
	cn <- 20

	fmt.Printf("cn: %v\n", cn)

	x := <-cn

	fmt.Printf("x1: %v\n", x)
	x = <-cn
	fmt.Printf("x2: %v\n", x)
}

//对比test1来看
func test2() {
	//定义一个通道
	var cn chan int
	fmt.Printf("cn: %v\n", cn)

	//和test1一样缓冲区为1
	cn = make(chan int, 1)
	fmt.Printf("cn 容量: %v\n", cap(cn))

	//下面代码不会报错，因为虽然缓冲区同样是1，但是通道在再次接收值之前，将缓冲区的值发送了出去，清空了缓冲区
	cn <- 10
	x := <-cn
	fmt.Printf("x1: %v\n", x)

	//因此通道可以再次接收值
	cn <- 20
	x = <-cn
	fmt.Printf("x2: %v\n", x)
}

//无缓冲通道
func test3() {
	//创建一个无缓冲通道
	cn := make(chan int)
	fmt.Printf("cn: %v\n", cn)
	//下面代码运行会报错，`deadlock`，因为通道没有缓冲区，再没有接收方接收通道值的时候，无法向通道发送值
	cn <- 10

}

//无缓冲通道2，对比test3来看
func test4() {
	//创建一个无缓冲通道
	cn := make(chan int)
	fmt.Printf("cn: %v\n", cn)
	//为了防止test3中的错误，我们可以创建一个`goroutine`接收通道的值
	go receive(cn)
	cn <- 10
	//向无缓冲通道发送值会导致阻塞，直到有另一个`goroutine`进行接收操作，这时无缓冲通道才会解除阻塞，继续执行
	//使用无缓冲通道可以将发送和接收同步化，因此无缓冲通道也叫`同步通道`

}

func receive(c chan int) {
	x := <-c
	fmt.Printf("x: %v\n", x)
}

func test5() {
	//创建一个缓冲区3的通道
	cn := make(chan int, 3)

	//通道接收两个值
	cn <- 10
	cn <- 20
	//然后关闭通道
	close(cn)

	//此时虽然缓冲区未满，但是通道已经关闭了，因此通道无法再接收值
	// cn <- 30

	//循环取出通道所有值，通道所有值被取出是会自动退出循环
	for {
		//从通道可以取出多个返回值，第一个为value，第二个为通道状态，false表示关闭，true表示开启
		//目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过len(ch)操作来判断通道是否被关闭。
		x, ok := <-cn
		fmt.Printf("x: %v\n", x)
		fmt.Printf("ok: %v\n", ok)
		if !ok {
			fmt.Printf("\"通道已关闭\": %v\n", "通道已关闭")
			break
		}
		fmt.Printf("\"通道仍开启\": %v\n", "通道仍开启")
	}
}

//多路复用，同时接收多个chan通道的值
func test6() {

	ch := make(chan int, 1)
	count := 10

	for i := 0; i < count; i++ {
		select {
		case x := <-ch:
			{
				fmt.Printf("x: %v\n", x)
			}

		case ch <- i:
			{

			}

		}
	}
}
