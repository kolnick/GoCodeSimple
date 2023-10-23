package main

import (
	"fmt"
	"reflect"
	"testing"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Age  int    `info:"age"  sex:"性别"`
}

func findTag(str interface{}) {
	elems := reflect.TypeOf(str).Elem()

	for i := 0; i < elems.NumField(); i++ {
		tagStr := elems.Field(i).Tag.Get("info")
		fmt.Println("info:", tagStr)
	}
}

func TestStruct(t *testing.T) {
	r := resume{"ccj", 18}
	findTag(&r)
}

type User struct {
	Username string `json:"username" validate:"required,max=50"`
	Email    string `json:"email" validate:"email"`
}

func TestGetAllFlags(d *testing.T) {
	user := User{
		Username: "john_doe",
		Email:    "john_doe@example.com",
	}
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, tag)
	}
}
