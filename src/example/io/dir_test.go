package io

import (
	"fmt"
	"os"
	"testing"
)

func TestDir(t *testing.T) {

	fileName := "./test1"
	err := os.Mkdir(fileName, os.ModePerm)
	if err != nil {
		fmt.Println("err ", err.Error())
	} else {
		fmt.Printf("%s 目录创建成功", fileName)
	}

	filePath2 := "./test2/abc/xyz"

	err = os.MkdirAll(filePath2, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 目录创建成功", filePath2)
	}
}
