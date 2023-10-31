package rpc

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	args := Args{7, 8}

	var reply int

	// 调用远程方法 Add
	err = client.Call("Calculator.Add", args, &reply)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result of %d + %d is %d\n", args.A, args.B, reply)

	// 调用远程方法 Multiply
	err = client.Call("Calculator.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result of %d * %d is %d\n", args.A, args.B, reply)
}
