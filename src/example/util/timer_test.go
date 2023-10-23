package util

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	// 创建一个新的计时器，会在持续时间之后将当前时间发送到channel上
	timer1 := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	data := <-timer1.C // chan time.Time

	fmt.Printf("%T\n", timer1.C)
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}

func TestAfter(t *testing.T) {
	ch1 := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-ch1
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
