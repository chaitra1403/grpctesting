package main

import (
	"context"
	"testing"

	hello "github.com/chaitra1403/hello-grpc/grpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGreeterClient_SayHello(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock GreeterClient
	mockGreeterClient := hello.NewMockGreeterClient(ctrl)

	// Setup mock expectations
	mockGreeterClient.EXPECT().
		SayHello(gomock.Any(), gomock.Any()).
		Return(&hello.HelloReply{
			Message:        "Hello, Chaitra!",
			Acknowledgment: "Hello received: Chaitra",
		}, nil).
		Times(1)

	// Call the mocked method directly without trying to establish a real connection
	req := &hello.HelloRequest{Name: "Chaitra"}
	res, err := mockGreeterClient.SayHello(context.Background(), req)

	// Assertions to ensure expected behavior
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Chaitra!", res.Message)
	assert.Equal(t, "Hello received: Chaitra", res.Acknowledgment)
}
