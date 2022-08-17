package main

import (
	"fmt"
	"sync"
)

var x int

var wg sync.WaitGroup

var m sync.Mutex
var rwm sync.RWMutex

func main() {
	// test1()
	// test2()
	//加锁访问，test2中是只要访问x 就加锁，无论是读还是写，单有时候我们要求实现多读单写，iOS中通过栅栏加group可以实现
	test3()
}
func test1() {
	wg.Add(2)
	//开启两个协程修改同一个变量 x 导致一个协程操作数据覆盖另一个的操作，
	go add()
	go add()

	wg.Wait()
	//所以最终值 <= 10000，因为多次累加后的结果可能被覆盖掉
	fmt.Printf("x: %v\n", x)
}

//加锁访问 互斥锁
func test2() {
	//和test1一样，只不过加了锁
	wg.Add(2)
	//虽然开启两个协程修改同一个变量 x ，但是因为加了互斥锁，同一时间只能有一个协程访问x
	go add1()
	go add1()

	wg.Wait()
	//所以最终值 == 10000
	fmt.Printf("x: %v\n", x)
}

//多读单写，相比于互斥锁性能更高，因为，大多数情况下更关注的写操作，读操作并不修改资源
func test3() {
	wg.Add(2)
	add2()
	wg.Wait()
}

//异步操作同一份资源问题
func add() {

	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func add1() {

	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg.Done()
}

func add2() {
	go rwmWrite()
	go rwmread()
}

func rwmWrite() {
	for i := 0; i < 5000; i++ {
		rwm.Lock()
		x = x + 1
		rwm.Unlock()
	}
	wg.Done()
}

func rwmread() {
	for i := 0; i < 5000; i++ {
		rwm.RLock()
		fmt.Printf("x: %v\n", x)
		rwm.RUnlock()
	}
	wg.Done()
}
