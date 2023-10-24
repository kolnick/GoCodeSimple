package _struct

import "fmt"

type Student struct {
	Name string
	Age  int32
}

func NewStudent(name string, age int32) Student {
	return Student{name, age}
}

func (stu Student) GetName() string {
	return stu.Name
}

func (this Student) String() string {
	result := fmt.Sprintf("name:%s--------age:%d", this.Name, this.Age)
	return result
}
