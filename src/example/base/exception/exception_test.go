package exception

import (
	"errors"
	"fmt"
	"testing"
)

func TestException(t *testing.T) {
	test()
	fmt.Println("上面的除法操作成功")
	fmt.Println("正常执行下面的逻辑")

	err := testCustomError()
	if err != nil {
		fmt.Println("自定义错误", err)
	}
	fmt.Println("上面的除法操作成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() {
	// 利用defer+recover 来捕获错误
	defer func() {
		//调用recover内置函数,可以捕获错误
		err := recover()
		// 如果没有部或错误,返回值是零值,nil
		if nil != err {
			fmt.Println("错误已经捕获")
			fmt.Println("err是:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func testCustomError() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
