package convert

import (
	"fmt"
	"strconv"
	"testing"
)

func TestItoa(t *testing.T) {
	s := strconv.Itoa(199)
	fmt.Printf("%T ,%v,长度%d \n", s, s, len(s))
}

func TestFormatInt(t *testing.T) {
	s := strconv.FormatInt(-19968, 16)
	s = strconv.FormatInt(-40869, 16)
	fmt.Printf("%T ,%v,长度%d \n", s, s, len(s))
}

func TestFormatFloat(t *testing.T) {
	s := strconv.FormatFloat(3.1415926, 'g', -1, 64)
	fmt.Printf("%T ,%v,长度%d \n", s, s, len(s))
}
