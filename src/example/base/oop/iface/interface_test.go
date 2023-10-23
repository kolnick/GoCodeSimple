package iface

import "fmt"

type Handler interface {
	handler()
}

type IHandler struct {
}

type MysqlHandler struct {
	IHandler
}
type OracleHandler struct {
	IHandler
}

// 实现接口的方法
func (handler MysqlHandler) handler() {
	fmt.Println("mysql handler")
}

// 实现接口的方法
func (handler OracleHandler) handler() {
	fmt.Println("oracle handler")
}

func handler(s Handler) {
	s.handler()
	fmt.Println("oracle handler")
}

func main() {
	c1 := MysqlHandler{}
	c2 := OracleHandler{}

	handler(c1)
	handler(c2)

}
