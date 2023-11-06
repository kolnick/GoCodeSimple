package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6666")
	if nil != err {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	var limit = 100000
	for {
		for i := 0; i < limit; i++ {
			data := []byte("12345678")
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println("Error sending:", err)
				return
			}
		}
		time.Sleep(time.Second)
	}

}
