package exception

import (
	"fmt"
	"testing"
)

func testA() {
	fmt.Println("func TestA")
}

func testB() {
	panic("func testB: panic")
}

func testC() {
	fmt.Println("func TestC")
}

func TestPanic(t *testing.T) {

	testA()
	testB() //遇到异常退出
	testC()
}
