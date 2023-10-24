package protobuf

import (
	protos "GoCodeSimple/src/example/serialization/protobuf/protos"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

// go get google.golang.org/protobuf/cmd/protoc-gen-go
func TestProtobuf(t *testing.T) {
	m := map[string]string{}
	m["user2"] = "john"
	m["user1"] = "kolnick"
	p := &protos.User{
		Id:           1,
		Name:         "kolnick",
		Email:        "kolnick@163.com",
		PhoneNumbers: []string{"123", "456", "1231"},
		Attributes:   m,
	}

	// 将消息序列化为字节流
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	var user2 protos.User
	err = proto.Unmarshal(data, &user2)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// 使用反序列化后的消息
	fmt.Println("Id:", user2.Id)
	fmt.Println("Name:", user2.Name)
	fmt.Println("Email:", user2.Email)
	fmt.Println("Attribute:", user2.Attributes)
	fmt.Println("Phone:", user2.PhoneNumbers)

}
