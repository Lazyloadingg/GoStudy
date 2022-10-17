package main

import (
	"fmt"
	"net/rpc"
)

func main() {
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
