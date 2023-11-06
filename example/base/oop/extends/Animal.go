package main

import "fmt"

type Animal struct {
	Age    int32
	Weight float32
}

func (animal *Animal) action() {
	fmt.Println("行动")
}

func (animal *Animal) showInfo() {
	fmt.Printf("动物的年龄是 %v 动物的体重是 %v\n", animal.Age, animal.Weight)
}

type Cat struct {
	// cat 继承了Animal里所有的属性和方法
	Animal
}

func (c *Cat) scratch() {
	fmt.Println("我是小猫,我会挠人")
}

func main() {
	cat := &Cat{}
	// 如果结构体中就用结构体中，否则从嵌入的结构体
	cat.Age = 3
	cat.Weight = 10.6

	cat.action()
	cat.showInfo()
	cat.scratch()
}

// 对Cat 绑定特有的方法
