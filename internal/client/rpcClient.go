package client

import (
	"fmt"

	"google.golang.org/grpc"
)

// GrpcClient represents the grpc-client infra needed to communicate with other grpc services
type GrpcClient struct {
	name    string
	addr    string
	options []grpc.DialOption
}

// New creates a new GrpcClient
func New(name string, def string) *GrpcClient {
	options := []grpc.DialOption{}
	options = append(options, grpc.WithInsecure())
	client := GrpcClient{
		name:    name,
		addr:    def,
		options: options,
	}
	return &client
}

// Dial connects to the grpc server
func (c *GrpcClient) Dial() *grpc.ClientConn {
	connection, err := grpc.Dial(c.addr, c.options...)
	checkErr(err)

	fmt.Println("Connected to ", c.addr)
	return connection
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
