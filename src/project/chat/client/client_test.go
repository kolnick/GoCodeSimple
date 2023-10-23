package client

import (
	"flag"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("<<<<<<链接服务器失败")
		return
	}

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println("<<<<<<链接服务器成功")
	// 启动客户端的业务
	select {}
}
