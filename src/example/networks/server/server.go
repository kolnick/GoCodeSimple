package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("服务器启动了...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("启动失败%v", err)
		return
	}
	fmt.Printf("启动成功%v \n", listen)

	for {
		// 等待客户端连接
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Printf("等待客户端链接失败%v", err)
		} else {
			fmt.Printf("接收到客户端的连接 con=%v,接收到的客户端的信息%v \n", conn, conn.RemoteAddr().String())
		}
		// 准备一个协程处理客户端的请求
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 创建一个切片 ，准备:将读取的数据放入切片
		buf := make([]byte, 1024)
		// 从conn 连接中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 将读取内容在服务器端输出
		fmt.Println(string(buf[0:n]))
	}
}
