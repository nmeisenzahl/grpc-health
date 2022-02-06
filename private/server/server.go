package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type server struct {
	srv     *grpc.Server
	address string
}

func New(address string) (*server, error) {
	s := server{
		srv:     grpc.NewServer(),
		address: address,
	}
	grpc_health_v1.RegisterHealthServer(s.srv, &server{})
	reflection.Register(s.srv)
	return &s, nil
}

func (s *server) Start(context.Context) error {
	conn, err := net.Listen("tcp", s.address)
	if err != nil {
		fmt.Println("cannot listen on address: " + s.address)
		return err
	}
	fmt.Println("start listening on " + s.address)
	fmt.Printf("Happy health-checking!\n")
	return s.srv.Serve(conn)
}

func (s *server) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *server) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return errors.New("not implemented")
}
