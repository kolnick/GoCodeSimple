package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer funcA()
	funcB()
	defer funcC()
	fmt.Println("main over...")
}

func funcA() {
	fmt.Println("this is funcA")
}

func funcB() {
	fmt.Println("this is funcB")
}
func funcC() {
	fmt.Println("this is funcC")
}
