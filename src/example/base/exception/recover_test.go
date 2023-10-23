package exception

import (
	"fmt"
	"testing"
)

func funA() {
	fmt.Println("func TestA")
}
func funB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复 获取recover的返回值", msg)
		}
	}()
	fmt.Println("this is funcB")

	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("funcB ")
		}
	}
}

func funcC() {
	defer func() {
		fmt.Println("执行延迟函数")
		msg := recover()
		fmt.Println("获取recover的返回值:", msg)

	}()

	fmt.Println("this is funcC")

	panic("funcC warning")

}

func TestRecover(t *testing.T) {
	funA()
	funB()
	funcC()
	fmt.Println("main over")
}
