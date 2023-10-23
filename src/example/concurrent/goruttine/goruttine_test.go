package goruttine

import (
	"fmt"
	"testing"
	"time"
)

func TestGo(t *testing.T) {

	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func TestMultiGo(t *testing.T) {
	go printNum()
	go printLetter()

	time.Sleep(3 * time.Second)
	fmt.Println("main over")
}

func printNum() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d \n", i)
	}
}

func printLetter() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c \n", i)
	}
}
