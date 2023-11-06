package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"-"` // 不解析
}

func TestStructJson(t *testing.T) {
	// Serialize int
	var data = &Person{1, "ccj", 5}
	// 序列化成字节
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Can't serislize", bytes)
	}
	fmt.Printf("%v => %v, '%v'\n", data, bytes, string(bytes))
	// Deserialize
	// 反序列化成对象
	var person Person
	err = json.Unmarshal(bytes, &person)
	if err != nil {
		fmt.Println("Can't deserislize", bytes)
	}
	fmt.Printf("%v => %v\n", bytes, person)
}

func TestMapJson(t *testing.T) {
	m := map[string][]string{
		"level":   {"debug"},
		"message": {"File not found", "Stack overflow"},
	}
	// 方便阅读的json格式
	if data, err := json.MarshalIndent(m, "", " "); err == nil {
		fmt.Printf("%s\n", data)
	}
}
