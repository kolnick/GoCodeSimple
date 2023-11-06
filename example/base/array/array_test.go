package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := [4]float64{67, 7, 89.8, 21}
	b := [...]int{2, 3, 5}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}

	fmt.Println()
	/**
	  遍历
	*/
	for value := range b {
		fmt.Print(value, "\t")
	}
	fmt.Println()

}

func second() {
	a := [4]int{1, 2, 3, 4}

	fmt.Printf("1. 变量a的内存地址是 %p 值为 %v \n\n", &a, a)
	//fmt.Printf("数组型变量a内存地址是:%p\n\n", a)
	//传值
	changeArrayVal(a)
	fmt.Printf("2.changeArrayVal 调用后:变量a的内存地址是%p 值为:%v \n\n", &a, a)
	//传引用
	changeArrayPtr(&a)
	fmt.Printf("3.changeArrayPtr 调用后:变量a的内存地址是%p 值为:%v \n\n", &a, a)

}

func changeArrayVal(a [4]int) {
	fmt.Printf("---------changeArrayVal函数内:值参数a的内存地址是%p,值为 %v \n", &a, a)
	//fmt.Printf("---------changeArrayPtr 函数内:值参数a的内存地址是:%p \n", a)
	a[0] = 99
}
func changeArrayPtr(a *[4]int) {
	fmt.Printf("---------changeArrayPtr函数内:指针参数a的内存地址是%p,值为 %v \n", &a, a)
	(*a)[1] = 250
}
