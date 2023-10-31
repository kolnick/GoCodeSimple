package grpc

import (
	pb "GoCodeSimple/src/framework/example/grpc/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type server struct {
	pb.CalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := req.A + req.B
	return &pb.Response{Result: result}, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

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

	client := pb.NewCalculatorClient(conn)

	req := &pb.Request{A: 10, B: 20}
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
