package main

import (
	"fmt"
	"net/rpc"
)

// 定义一个Helloservice客户端结构体，方便调用
type HelloServiceClient struct {
	*rpc.Client
}

// 同样定义服务名称
const HelloServiceName = "/server/main/HelloService"

// 拨号rpc服务，并返回client
func DialHelloService(net, addr string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(net, addr)
	if err != nil {
		fmt.Printf("\"拨号失败\": %v\n", "拨号失败")
		return nil, err
	}

	return &HelloServiceClient{c}, nil

}

// 调用Hello方法
func (c *HelloServiceClient) Hello(param Param, reply *int) error {
	return c.Client.Call(HelloServiceName+".Hello", param, reply)
}
