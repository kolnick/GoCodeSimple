package grpc

import (
	"GoCodeSimple/framework/example/grpc/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type server struct {
	__.CalculatorServer
}

func (s *server) Add(ctx context.Context, req *__.Request) (*__.Response, error) {
	result := req.A + req.B
	return &__.Response{Result: result}, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	__.RegisterCalculatorServer(s, &server{})

	fmt.Println("Server is running on port 1234")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := __.NewCalculatorClient(conn)

	req := &__.Request{A: 10, B: 20}
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
