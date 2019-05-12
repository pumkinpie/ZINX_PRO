package main

import (
	"ZINX_PRO/zinx/ziface"
	"ZINX_PRO/zinx/znet"
	"fmt"
)

/*
	基于 Zinx框架来开发，服务端相应程序
*/

// PingRouter test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")

	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		", data = ", string(request.GetData()))

	// 发送消息给客户端
	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

// HelloRouter test 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

// test Handle
func (pr *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")

	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		", data = ", string(request.GetData()))

	// 发送消息给客户端
	err := request.GetConnection().SendMsg(1, []byte("bee...bee...bee"))
	if err != nil {
		fmt.Println(err)
	}
}

//创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}

	// 设置两个链接属性，在连接创建之后
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "Aceld")
	conn.SetProperty("Home", "https://www.jianshu.com/u/35261429b7f1")
}

//连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("DoConneciotnLost is Called ... ")

	// 在连接销毁之前，查询conn的Name，Home属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}

	fmt.Println("DoConnectionLost is Called ... ")
}

func main() {
	// 1 创建一个server句柄，使用Zinx的 api
	s := znet.NewServer("[Zinx V0.9]")
	// 2 注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	// 3 给当前的zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	// 4 启动server
	s.Serve()
}
