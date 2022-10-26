package main

import (
	"fmt"
	"net/http"
	"reflect"
)

// 自定义方法类型
type CusFunc func(a string, b string)

// 定义接口
type CusInterface interface {
	Hello(a, b string)
}

// 方法也可以实现接口
func (c CusFunc) Hello(a, b string) {
	fmt.Printf("\"hello\": %v\n", "hello")

}

func test(fn CusInterface) {
	// fn("666", "777")
	fmt.Printf("reflect.ValueOf(fn).String(): %v || %v\n", reflect.TypeOf(fn), reflect.ValueOf(fn))
	// fn.Hello(f)

}
func initServer() {

	test(CusFunc(func(a, b string) {
		fmt.Printf("(a + b): %v\n", (a + b))
	}))

	//第二个参数为Handler类型
	// http.HandleFunc("/", hello)

	http.Handle("/", timeMiddle(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("\"服务启动失败\": %v\n", "服务启动失败")
	}
	fmt.Printf("\"服务启动成功\": %v\n", "服务启动成功")

	//1. 给一个路由添加中间件其实就是将这个路由包装一层
	//2. 为了和原注册路由方法保持一致，用一个函数封装路由，并返回和路由相同的类型

}

// 服务函数
func hello(w http.ResponseWriter, r *http.Request) {

	resS := "hello world1"
	fmt.Printf("resS: %v\n", resS)
	res := []byte(resS)
	w.Write(res)

}

// 自定义中间件，`hello`函数本质是一个`http.Handler`类型，这里同样返回`http.Handler`类型
// 参数就是需要添加中间件的函数，同样是`http.Handler`类型
func timeMiddle(fn http.Handler) http.Handler {
	//构造一个`http.Handler`类型并返回,HandlerFunc实现了`http.Handler`接口，因此属于`http.Handler`类型
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\"调用前\": %v\n", "调用前")
		fn.ServeHTTP(w, r)
		fmt.Printf("\"调用后\": %v\n", "调用后")
	})
}
