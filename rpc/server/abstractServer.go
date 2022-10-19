package main

import "net/rpc"

// 定义方法名，为了防止重名，最好名字里带上所属的包，这里自己定义规则
const HelloServiceName = "/server/main/HelloService"

// 抽离服务端方法为接口，哪个`具体服务`想被rpc客户端调用，就自己实现接口方法，方便扩展和迁移
type HelloServiceInterface interface {
	Hello(param Param, reply *int) error
}

// 定义一个实现方法的结构体
type HelloService struct{}

// HelloService实现`HelloServiceInterface`接口
func (h *HelloService) Hello(param Param, reply *int) error {
	*reply = param.Width * param.Height
	return nil
}

// 为了防止注册服务容易写错问题，将HelloServiceInterface服务注册抽离成方法
func RegisterHelloService(s HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, s)
}
