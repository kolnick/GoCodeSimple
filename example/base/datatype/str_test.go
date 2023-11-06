package datatype

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestStr(t *testing.T) {
	// 遍历字符串
	str := "hello world"
	for i, value := range str {
		fmt.Printf("第 %d 位的ASCII的值是%d 字符%c \n", i, value, value)
	}
}
func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
