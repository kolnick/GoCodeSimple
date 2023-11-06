package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNoBufferChannel(t *testing.T) {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行...")
		// 将666 发送给c
		c <- 666
	}()

	// 从c中接受数据,并赋值给num 执行到的时候就进行等待阻塞
	num := <-c

	fmt.Println("num=", num)
	fmt.Println("main goroutine 结束...")
}

func TestBufferChannel(t *testing.T) {
	// 创建带有缓冲的channel
	c := make(chan int, 3)

	fmt.Println("len(c)=", len(c), ",cap(c)", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程序运行,发送的元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main 结束")
}

func TestCloseChannel(t *testing.T) {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close(c)
	}()

	for {
		// ok为true表示channel没有关闭，false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("Main Finished...")
}
func TestChannelRange(t *testing.T) {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			// 阻塞接收数据
			c <- i
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("Main Finished...")
}

func TestChannelSelect(t *testing.T) {

	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	testFunc(c, quit)
}

func testFunc(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
