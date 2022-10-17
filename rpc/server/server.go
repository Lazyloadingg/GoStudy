package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// 定义服务
type TestService struct{}

func main() {
	test := new(TestService)
	//注册服务
	rpc.Register(test)
	//绑定协议
	rpc.HandleHTTP()
	//监听服务
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		fmt.Printf("\"rpc监听失败\": %v\n", "rpc监听失败")
		panic(err)
	}
	fmt.Printf("\"rpc监听成功\": %v\n", "rpc监听成功")

}

// rpc函数（必须有接收者，指定属于哪个服务）
// 函数名首字母必须大写(公开函数)
// 第一个参数：入参
// 第二个参数：出参，且必须为指针类型
// 返回值必须有，且为error
func (s *TestService) Area(param map[string]int, res *int) error {

	fmt.Printf("\"方法调用\": %v\n", "方法调用")
	*res = param["width"] * param["height"]
	return nil
}
