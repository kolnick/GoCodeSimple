package rpc

import (
	"net"
	"net/rpc"
	"testing"
)

// 定义一个计算器服务
type Calculator struct{}

// 定义两个整数相加的方法
func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// 定义两个整数相乘的方法
func (c *Calculator) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 定义参数结构体
type Args struct {
	A, B int
}

func TestServer(t *testing.T) {
	calculator := new(Calculator)

	// 注册服务
	err := rpc.Register(calculator)
	if err != nil {
		panic(err)
	}

	// 创建 TCP 服务器
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
