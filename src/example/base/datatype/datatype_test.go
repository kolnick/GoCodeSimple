package datatype

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConst(t *testing.T) {
	const (
		zero       = 0.0
		size int64 = 1024
		eof        = -1
	)
	const u, v float32 = 0, 3

	const a, b, c = 3, 4, "foo"

	const PI float64 = 3.141592653

	const mask = 1 << 3
}

func TestBool(t *testing.T) {
	var bol = false
	println(bol)

	b := bool(false)
	println(b)
}

func TestChar(t *testing.T) {
	var ch = byte('a')
	println(ch)
}
func TestString(t *testing.T) {

	// 字符串遍历
	str := "golang你好"

	for i, value := range str {
		fmt.Printf("索引为:%d, 具体的值 %c\n", i, value)
	}
	fmt.Println("--------")
	// 使用rune 遍历
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}
	// 字符串转整数
	num1, _ := strconv.Atoi("1234")
	fmt.Println(num1)
	// 整数转字符串
	str1 := strconv.Itoa(num1)
	fmt.Println(str1)
	// 包含子串的数量
	count := strings.Count("123456123123132", "123")
	fmt.Println("包含子串的数量", count)
	// 不区分大小写字符串比较
	equals := strings.EqualFold("wqwe", "WQWe")
	fmt.Println("是否相同:", equals)
	fmt.Println("是否相同", "hello" == "HELLO")

	// 返回子串在字符串第一次出现的索引值
	firstIndex := strings.Index("golangstartggggggstart", "start")
	fmt.Println("第一次出现索引的位置", firstIndex)
	// 字符串替换 n表示替换的个数 -1表示全部替换
	replace := strings.Replace("gogogo", "go", "java", 2)
	fmt.Println("替换后的字符串:", replace)

	// 字符串切割
	strArr := strings.Split("go-java-python", "-")
	fmt.Println(strArr)

	// 字符串大小写切换
	lowStr := strings.ToLower("abcSs")
	fmt.Println("切换成小写", lowStr)
	upStr := strings.ToUpper(lowStr)
	fmt.Println("切换成大写", upStr)
	// 字符串左右两边的空格
	trimStr := strings.TrimSpace("    go lang java     ")
	fmt.Println("去除左右两边的空格", trimStr)

	trimfe := strings.Trim("~golang~", "~")
	fmt.Println("去除左右两边的指定字符", trimfe)

	hasPre := strings.HasPrefix("你好的啊", "你好")
	fmt.Println("是否是XX开头", hasPre)

	suffix := strings.HasSuffix("你好的啊", "你好")
	fmt.Println("是否是XX结尾", suffix)
}

func TestArray(t *testing.T) {

}

func TestMap(t *testing.T) {
	var m map[string]int //key=string value=int
	print(m)
}

func TestFloat(t *testing.T) {

	f := float32(32)
	println(f)

	f1 := 1.1
	println(f1)

	f2 := float64(1.1)

	println(f2)

	num := 1.1
	fmt.Printf("数据类型是 %T", num)
}

func TestNew(t *testing.T) {
	// new 分配内存,new函数的实参是一个类型而不是具体数值哦,new函数返回值是对应类型的指针

	num := new(int)
	fmt.Printf("num的类型:%T ,num的值是%v ,地址是:%v,指针指向的值是%v", num, num, &num, *num)
}
