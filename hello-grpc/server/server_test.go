package main

import (
	"context"
	"testing"

	hello "github.com/chaitra1403/hello-grpc/proto"
	mock "github.com/chaitra1403/hello-grpc/proto"
	"github.com/stretchr/testify/assert"
)

// Mock implementation of GreeterServer
type mockGreeterServer struct {
	hello.UnimplementedGreeterServer
}

// SayHello implements the GreeterServer interface
func (m *mockGreeterServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{
		Message:        "Hello " + req.Name,
		Acknowledgment: "Hello received: " + req.Name,
	}, nil
}

// TestGreeterServer_SayHello tests the gRPC server
func TestGreeterServer_SayHello(t *testing.T) {
	server := &mockGreeterServer{}

	// Create a context and request
	ctx := context.Background()
	req := &mock.HelloRequest{Name: "Chaitra"}

	// Call the SayHello method on the server
	resp, err := server.SayHello(ctx, req)

	// Assert no error and response correctness
	assert.NoError(t, err)
	assert.Equal(t, "Hello Chaitra", resp.Message)
	assert.Equal(t, "Hello received: Chaitra", resp.Acknowledgment)
}
