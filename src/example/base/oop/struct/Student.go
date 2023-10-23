package _struct

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
