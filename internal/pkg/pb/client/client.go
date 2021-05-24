package client

import (
	"log"
	"ports/internal/pkg/pb"

	"google.golang.org/grpc"
)

// PortsServiceClient wrapper for GRPC client to ports service
type PortsServiceClient struct {
	Client pb.PortsServiceClient
}

func (c *PortsServiceClient) Connect() {
	// TODO handle disconnect
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.Client = pb.NewPortsServiceClient(conn)
}
