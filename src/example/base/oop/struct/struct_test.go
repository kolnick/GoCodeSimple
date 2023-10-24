package _struct

import (
	"fmt"
	"testing"
)

func TestAnonymousStruct(t *testing.T) {
	anonymous := struct {
		id   int
		name string
		age  int
	}{id: 1, name: "kolnick", age: 21}

	fmt.Println(anonymous)
}

func getFlower() (f Student) {
	s := Student{"kolnick", 21}
	return s
}

type User struct {
	string
	byte
	int8
	double
}

type double float64

func TestAnonymousField(t *testing.T) {

	user := User{"kolnick", 'm', 35, 177.5}
	fmt.Println(user)
}

func TestPrint(t *testing.T) {
	student := Student{"kolnick", 21}
	fmt.Println(student)
}
