package file

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	path := "../../../res/微信截图_20230615134853.png"
	printMessage(path)
}

func printMessage(filepath string) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("数据类型:%T\n", fileInfo)
		fmt.Println("文件名:", fileInfo.Name())
		fmt.Println("是否为目录:", fileInfo.IsDir())
		fmt.Println("文件大小:", fileInfo.Size())
		fmt.Println("文件权限:", fileInfo.Mode())
		fmt.Println("文件最后修改时间:", fileInfo.ModTime())
	}
}

func TestCreateFile(t *testing.T) {
	filename := "../test1/a.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
	} else {
		fmt.Println("文件创建成功", file)
	}
}

func TestOpenFile(t *testing.T) {
	filename := "../test1/a.txt"
	// 打开文件
	file2, err := os.Open(filename)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功 %v\n", filename, file2)
	}
	defer file2.Close()

	filName4 := "../test1/abc2.txt"
	// 以读写方式打开  如果文件不存在就创建
	file4, err := os.OpenFile(filName4, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功 %v \n", filName4, file4)
	}
	defer file4.Close()
}

// 可用于删除文件或者删除目录
func TestRemoveFile(t *testing.T) {
	fileName := "../test2/"

	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s 删除成功", fileName)
	}

	err = os.RemoveAll(fileName)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("%s 删除成功", fileName)
	}
}

func TestReadFile(t *testing.T) {
	fileName := "../test1/1.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("文件读取错误")
	}
	// 将读取到的字节转换成字符串
	fmt.Println(string(file))
}

func TestReadFileWithBuffer(t *testing.T) {
	list := readFileWithBuffer("../test1/1.txt")
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

func readFileWithBuffer(path string) *list.List {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	list := list.New()
	for {
		str, err := reader.ReadString('\n')
		list.PushBack(str)
		if err == io.EOF { // io.EOF 表示已经到文件结尾
			break
		}
	}
	return list
}

func TestWriteFile(t *testing.T) {
	fileName := "../test1/1.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件异常", err.Error())
	} else {
		n, err := file.Write([]byte("abcde12345"))
		if err != nil {
			fmt.Println("写入文件异常", err.Error())
		} else {
			fmt.Println("写入ok", n)
		}
		n, err = file.WriteString("中国字")
		if err != nil {
			fmt.Println("写入文件异常", err.Error())
		} else {
			fmt.Println("写入ok", n)
		}
	}
}

func TestCopyFile(t *testing.T) {
	// 源文件相对路径
	srcFile := "../test1/1.txt"
	// 准备生成的目标文件路径
	destFile := "../test1/2.txt"

	total, err := copyFile(srcFile, destFile)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("复制OK", total)
	}
}

func copyFile(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
