package main

import (
	"fmt"
)

func main() {
	// 方式1
	myArray := []int{1, 2, 3, 4}
	printArray(myArray)
	fmt.Printf("%v", len(myArray))

	// 开辟3个空间
	myArray2 := make([]int, 3, 5)
	for index, _ := range myArray2 {
		myArray2[index] = 1
	}

	// 追加元素
	myArray2 = append(myArray2, 5)
	printArray(myArray2)

	// 切片截取
	fmt.Println("number 0:3", myArray2[0:3])
	// 截取方式2
	fmt.Println("number 0:2", myArray2[:2])

	// copy数组
	copyArray := make([]int, 5)
	copy(copyArray, myArray2)
	printArray(copyArray)
}

// 引用传递
func printArray(array []int) {
	for index, value := range array {
		fmt.Println("index=", index, "value=", value)
	}
}
