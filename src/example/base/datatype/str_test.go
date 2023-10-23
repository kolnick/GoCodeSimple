package datatype

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	// 遍历字符串
	str := "hello world"
	for i, value := range str {
		fmt.Printf("第 %d 位的ASCII的值是%d 字符%c \n", i, value, value)
	}

}
