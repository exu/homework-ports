package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"ports/internal/pkg/pb"
	"ports/internal/pkg/storage/mongo"

	"google.golang.org/grpc"
)

func NewPortsServiceServer() *PortsServiceServer {
	portsRepository, err := mongo.NewPortsRepository()
	if err != nil {
		log.Fatalf("Handle this error correctly: %v", err)
	}
	return &PortsServiceServer{
		Repo: portsRepository,
	}
}

type PortsServiceServer struct {
	Repo mongo.PortsRepository
	pb.UnimplementedPortsServiceServer
}

func (s PortsServiceServer) Insert(context.Context, *pb.Port) (*pb.Port, error) { return nil, nil }
func (s PortsServiceServer) Get(context.Context, *pb.Code) (*pb.Port, error)    { return nil, nil }
func (s PortsServiceServer) List(context.Context, *pb.Codes) (*pb.Ports, error) { return nil, nil }
func (s PortsServiceServer) Delete(context.Context, *pb.Code) (*pb.Port, error) { return nil, nil }

func Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", "9091"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPortsServiceServer(grpcServer, NewPortsServiceServer())
	grpcServer.Serve(lis)
}
