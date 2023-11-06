package error

import (
	"errors"
	"fmt"
	"math"
	"os"
	"testing"
)

func TestNewError(t *testing.T) {

	err := errors.New("自己创建的错误")
	fmt.Println(err.Error())
	fmt.Printf("%T\n", err)
	fmt.Println("-----------")

	err2 := fmt.Errorf("错误的类型%d", 10)
	fmt.Println(err2.Error())
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)
	fmt.Println("-----------")

	age, err3 := checkAge(-12)
	if err3 != nil {
		fmt.Println(err3.Error())
		fmt.Println(err3)
	} else {
		fmt.Println(age)
	}
}

func checkAge(age int) (string, error) {
	if age < 0 {
		err := fmt.Errorf("输入的年龄是%d,数值是负数,有错误", age)
		return "", err
	} else {
		return fmt.Sprintf("输入年龄是%d ", age), nil
	}

}

func TestError(t *testing.T) {

	res := math.Sqrt(-100)
	fmt.Println(res)
}

func TestError2(t *testing.T) {

	open, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(open.Name())
	}
}
