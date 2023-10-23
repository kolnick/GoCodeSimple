package iface

import (
	"fmt"
	"testing"
)

type A interface {
}

type Cat struct {
	name string
	age  int
}
type Person struct {
	name string
	sex  string
}

func TestEmptyTest(t *testing.T) {
	var a1 A = Cat{"aaa", 19}
	showInfo(a1)
}

func showInfo(a A) {
	fmt.Printf("%T %v\n", a, a)
}

func TestAnyValueMap(t *testing.T) {
	map1 := make(map[string]interface{})
	map1["name"] = "kolnick"
	map1["age"] = 13
	fmt.Println(map1)
}

func TestSlice(t *testing.T) {
	slice := make([]interface{}, 0, 10)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
}
