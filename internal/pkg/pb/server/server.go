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

func (s PortsServiceServer) Insert(ctx context.Context, port *pb.Port) (*pb.Port, error) {
	err := s.Repo.Save(port)
	return port, err
}
func (s PortsServiceServer) Get(ctx context.Context, code *pb.Code) (*pb.Port, error) {
	port, exists := s.Repo.Get(code.Code)
	if !exists {
		return port, fmt.Errorf("Can't find port with code:%s", code.Code)
	}

	return port, nil
}
func (s PortsServiceServer) List(ctx context.Context, codes *pb.Codes) (ports *pb.Ports, err error) {
	dbPorts, err := s.Repo.List(codes.Code...)
	if err != nil {
		return nil, err
	}
	ports.Ports = dbPorts
	return
}
func (s PortsServiceServer) Delete(ctx context.Context, code *pb.Code) (port *pb.Port, err error) {
	err = s.Repo.Delete(code.Code)
	return port, err
}

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
