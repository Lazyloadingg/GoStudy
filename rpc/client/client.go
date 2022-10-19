package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Param struct {
	Width, Height int
}

func main() {
	// gorpcClient()
	jsonrpcClient()
	// absClient()
}

// rpc客户端
func gorpcClient() {
	//连接远程rpc
	client, err := rpc.DialHTTP("tcp", ":9900")
	if err != nil {
		fmt.Printf("\"连接rpc服务失败\": %v\n", "连接rpc服务失败")
	}
	//定义出参
	var res int
	//调用远程服务 args: (服务名.方法名) ，入参，出参，返回值
	err = client.Call("TestService.Area", map[string]int{"width": 5, "height": 7}, &res)
	if err != nil {
		fmt.Printf("\"TestService.Area rpc调用失败\": %v--%v\n", "TestService.Area rpc调用失败", err)
		return
	}
	fmt.Printf("TestService.Area rpc调用成功res: %v\n", res)
}

// jsonrpc客户端
func jsonrpcClient() {
	client, err := jsonrpc.Dial("tcp", ":9900")
	if err != nil {
		fmt.Printf("\"连接rpc服务失败\": %v\n", "连接rpc服务失败")
	}
	//定义出参
	var res int
	//同步调用
	//调用远程服务 args: (服务名.方法名) ，入参，出参，返回值
	err = client.Call("TestService.Perimeter", Param{Width: 2, Height: 4}, &res)
	if err != nil {
		fmt.Printf("\"TestService.Perimeter rpc调用失败\": %v--%v\n", "TestService.Perimeter rpc调用失败", err)
		return
	}
	fmt.Printf("TestService.Perimeter rpc调用成功res: %v\n", res)

	//异步调用
	cn := make(chan *rpc.Call, 1)
	client.Go("TestService.Perimeter", Param{Width: 2, Height: 4}, &res, cn)
	x := <-cn
	fmt.Printf("x: %v--%v\n", x, res)

}

// 封装客户端调用
func absClient() {
	client, err := DialHelloService("tcp", ":9900")
	if err != nil {
		fmt.Printf("\"连接rpc服务失败\": %v\n", "连接rpc服务失败")
	}
	var reply int
	client.Hello(Param{Width: 2, Height: 4}, &reply)
	fmt.Printf("reply: %v\n", reply)
}
