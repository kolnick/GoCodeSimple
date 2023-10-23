package datatype

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	// 生成重复3次
	repeat := strings.Repeat("hello ", 3)
	fmt.Println(repeat)
}
