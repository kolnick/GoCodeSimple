package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file1 := "./test1/1.txt"
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Println("文件打开异常", err.Error())
	} else {
		fmt.Println(string(data))
	}

	fileName2 := "./test1/1.txt"
	s1 := "测试"
	err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
	if err != nil {
		fmt.Println("写入文件异常", err.Error())
	} else {
		fmt.Println("文件写入ok")
	}

	// 文件复制
	err = ioutil.WriteFile(fileName2, data, os.ModePerm)
	if err != nil {
		fmt.Println("文件复制异常", err.Error())
	} else {
		fmt.Println("文件复制成功")
	}

	// 遍历目录
	dirName := "./"
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("目录遍历异常", err.Error())
	} else {
		for i, v := range fileInfos {
			fmt.Println(i, v.Name(), v.IsDir(), v.Size(), v.ModTime())
		}
	}

	// 创建目录
	filename, err := ioutil.TempDir("./", "temp")
	if err != nil {
		fmt.Println("创建目录失败", err.Error())
	} else {
		fmt.Println(filename)
	}

	// 创建文件
	file, err := ioutil.TempFile(filename, "temp")
	if err != nil {
		fmt.Println("创建文件失败", err.Error())
	} else {
		file.WriteString("写入内容:" + file.Name())
	}
	defer file.Close()

}
