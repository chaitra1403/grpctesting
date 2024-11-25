package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/chaitra1403/hello-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement hello.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Print a message each time a request is received
	fmt.Println("Received request from:", in.Name)

	// Server acknowledges that the "Hello" was received
	ackMessage := "Hello received: " + in.Name
	return &pb.HelloReply{
		Message:        "Hello " + in.Name,
		Acknowledgment: ackMessage,
	}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register Greeter server
	pb.RegisterGreeterServer(grpcServer, &server{})

	// Start the server
	fmt.Println("Starting server on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}
