package convert

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrToInt(t *testing.T) {
	// 字符串转换成int类型
	value, _ := strconv.Atoi("100")
	fmt.Printf("%T------%v", value, value+2)
}

func TestStrToFloat(t *testing.T) {
	// 将字符串转换成float类型
	pi := "3.1415926"
	num, _ := strconv.ParseFloat(pi, 64)
	fmt.Printf("%T------%v\n", num, num)
}
func TestStrToBool(t *testing.T) {
	// 将字符串转换成float类型
	str := "True"
	num, _ := strconv.ParseBool(str)
	fmt.Printf("%T------%v\n", num, num)
}
