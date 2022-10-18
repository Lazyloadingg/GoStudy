package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义服务
type TestService struct{}

type Param struct {
	Width  int
	Height int
}

func main() {

	//1. 创建服务
	//2. 注册服务

	// gorpc()
	jsonrpcServer()
}

// gorpc，只能go语言服务之间调用，不能跨语言
func gorpc() {
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

func (s *TestService) Perimeter(args Param, reply *int) error {
	*reply = (args.Height + args.Width) * 2
	return nil
}

// 可以跨语言
func jsonrpcServer() {
	err := rpc.Register(new(TestService))
	if err != nil {
		fmt.Printf("\"jsonrpc注册失败\": %v\n", "jsonrpc注册失败")
		return
	}
	lis, err := net.Listen("tcp", ":9900")

	if err != nil {
		fmt.Printf("\"jsonrpc监听失败\": %v\n", "jsonrpc监听失败")
		return
	}

	for {
		//接收连接
		con, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Printf("客户端进入")
			jsonrpc.ServeConn(conn)
		}(con)
	}
}
