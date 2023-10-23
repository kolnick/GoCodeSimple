package function

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	i := add(1, 2)
	println(i)
}

/*函数的定义*/
func add(s1 int32, s2 int32) int32 {
	return s1 + s2
}

// 返回多个值

func foo(s1 int32, s2 int32) (r1 int32, r2 int32) {
	return s1 + s2, s1 - s2
}

func TestTwoValueReturn(t *testing.T) {
	r1, r2 := foo(1, 2)
	println(r1)
	println(r2)
}

func TestAnonymousFunc(t *testing.T) {

	result := func(a, b float64) float64 {
		return a + b
	}(2, 3)

	fmt.Println(result)
}
