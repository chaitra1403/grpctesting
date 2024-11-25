package main

import (
	"context"
	"testing"

	mock "github.com/chaitra1403/hello-grpc"
	hello "github.com/chaitra1403/hello-grpc/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGreeterClient_SayHello tests the gRPC client
func TestGreeterClient_SayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock GreeterClient
	mockGreeterClient := mock.NewMockGreeterClient(ctrl)

	// Set up mock expectations
	mockGreeterClient.EXPECT().
		SayHello(gomock.Any(), gomock.Any()).
		Return(&hello.HelloReply{
			Message:        "Hello Chaitra",
			Acknowledgment: "Hello received: Chaitra",
		}, nil).
		Times(1)

	// Create a mock context and request
	ctx := context.Background()
	req := &hello.HelloRequest{Name: "Chaitra"}

	// Call the SayHello method on the mock client
	resp, err := mockGreeterClient.SayHello(ctx, req)

	// Assert no error and response correctness
	assert.NoError(t, err)
	assert.Equal(t, "Hello Chaitra", resp.Message)
	assert.Equal(t, "Hello received: Chaitra", resp.Acknowledgment)
}
