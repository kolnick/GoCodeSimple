package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

//实现了带缓冲的I/O读写 达到高效读写

func TestReader(t *testing.T) {
	fileName1 := "./test1/1.txt"

	file, _ := os.Open(fileName1)

	reader := bufio.NewReader(file)

	fmt.Printf("%T\n", reader)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Println(readString)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
	}
	defer file.Close()
}

func TestWriter(t *testing.T) {
	fileName1 := "./test1/1.txt"
	file2, _ := os.Open(fileName1)

	reader := bufio.NewReader(file2)

	fileName3 := "./test1/3.txt"
	file3, _ := os.OpenFile(fileName3, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	writer := bufio.NewWriter(file3)
	for {
		bytes, err := reader.ReadBytes(' ')
		writer.Write(bytes)
		writer.Flush()
		if err == io.EOF {
			fmt.Println("文件读取完毕!")
			break
		}
	}
	defer file2.Close()
	defer file3.Close()
}

func TestWrite2(t *testing.T) {
	path := "d:/test2.txt"
	writeFile(path)
}
func writeFile(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("\nhello world")
	// 流带缓冲区 只有
	writer.Flush()
	s := os.FileMode(0666).String()
	fmt.Println(s)
}
