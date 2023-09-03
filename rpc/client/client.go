package client

import "google.golang.org/grpc"

type RPCClient struct {
	grpc.ClientConn
	address string
}

func NewRPCClient(serverAddr string) {
	conn, err := grpc.Dial(serverAddr)
	if err != nil {
	}
	defer conn.Close()
}
