package oop

import (
	"fmt"
	"testing"
)

type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

func TestCopy(t *testing.T) {

	d1 := Dog{"dd", "black", 2, "two"}

	fmt.Printf("d1 :%T ,%v,%p \n", d1, d1, &d1)

	// 深拷贝
	d2 := d1
	fmt.Printf("d2 :%T ,%v,%p \n", d2, d2, &d2)

	d2.name = "maomao"

	fmt.Println("d2 修改后", d2)

	fmt.Println("d1 ", d1)

	fmt.Println("-----------")

	// 浅拷贝直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3 :%T ,%v,%p \n", d3, d3, &d3)
	d3.name = "qiuqiu"
	d3.color = "白色"
	d3.kind = "sala"

	fmt.Println("d3 修改后:", d3)
	fmt.Println("d1 :", d1)

	fmt.Println("-----------")
}
