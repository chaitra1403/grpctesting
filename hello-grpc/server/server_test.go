package main

import (
	"context"
	"testing"

	pb "github.com/chaitra1403/hello-grpc/grpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGreeterServer_SayHello(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create an instance of the actual server
	gs := &server{} // This is the server you defined in the original server.go file

	// Prepare the input request
	req := &pb.HelloRequest{Name: "Chaitra"}

	// Call the SayHello method on the server
	res, err := gs.SayHello(context.Background(), req)

	// Assert that there is no error
	assert.NoError(t, err)

	// Assert the expected response message and acknowledgment
	assert.Equal(t, "Hello Chaitra", res.Message)
	assert.Equal(t, "Hello received: Chaitra", res.Acknowledgment)
}
