package protobuf

import (
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

// go get google.golang.org/protobuf/cmd/protoc-gen-go
func TestProtobuf(t *testing.T) {
	p := &Person{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// 将消息序列化为字节流
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
}
