package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("客户端启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("启动客户端失败!")
		return
	}
	fmt.Println("客户端启动成功...", conn)

	//代表终端标准输入
	reader := bufio.NewReader(os.Stdin)
	// 从终端读取一行用户输入的信息
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败,err=", err)
		return
	}
	write, err2 := conn.Write([]byte(readString))
	if err2 != nil {
		fmt.Printf("连接失败 err", err2)
	}
	fmt.Printf("终端数据通过客户端发送成功,一共发送了%d字节的数据", write)
}
