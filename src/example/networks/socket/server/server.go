package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":6666")
	if nil != err {
		fmt.Println("Error creating listener:", err)
		return
	}
	fmt.Println("Listening on port 6666...")

	for {
		// 等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// 开始接收数据
		go receiveData(conn)
	}
}

func receiveData(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的数据
	buf := make([]byte, 1024)
	t0 := time.Now()

	// 第几次 接收到包  平均每秒处理多少包
	//averageSecRecv := list.New()

	// 总接收包
	totalRecvCount := 0
	// 第几轮时间接收包
	totalSecLoop := 0
	loopRecvCount := 0
	fmt.Println("------------")

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			return
		}
		totalRecvCount++
		loopRecvCount++
		t1 := time.Now()
		if t1.UnixMilli()-t0.UnixMilli() >= time.Second.Milliseconds() {
			t0 = t1
			totalSecLoop++
			fmt.Printf("第%v次处理%v个包\n", totalSecLoop, loopRecvCount)
			loopRecvCount = 0
			//averageSecRecv.PushFront(totalRecvCount)
		}
		// 将接收到的字节数据转换为字符串并打印
		//data := string(buf[:n])
		//fmt.Println("Received data:", data)

		//checkCount := 10
		//if loopRecvCount > checkCount && averageSecRecv.Len() > 0 {
		//	//10次内平均每秒发包多少次
		//	totalAvgTimeCount := 0
		//	for e := averageSecRecv.Front(); e != nil; e = e.Next() {
		//
		//		ct := e.Value.(int)
		//		fmt.Printf("ct:%d\n", ct)
		//		totalAvgTimeCount += ct
		//	}
		//	fmt.Printf("totalAvgTimeCount:%d\n", totalAvgTimeCount)
		//	fmt.Printf("在%v次中平均每秒接收处理%v个包\n", checkCount, totalAvgTimeCount/averageSecRecv.Len())
		//	loopRecvCount = 0
		//}
	}
}
