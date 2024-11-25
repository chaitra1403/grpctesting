package main

import (
	"context"
	"fmt"

	pb "github.com/chaitra1403/hello-grpc/grpc"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewGreeterClient(conn)

	// Call SayHello
	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		fmt.Println("Error calling SayHello:", err)
		return
	}

	// Print the response and acknowledgment
	fmt.Println("Server response:", res.Message)
	fmt.Println("Acknowledgment:", res.Acknowledgment)
}
