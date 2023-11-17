package server

import (
	"google.golang.org/grpc"
)

type GrpcServer struct {
	host       string
	port       string
	grpcServer *grpc.Server
}

func NewGrpcServer(host, port string) *GrpcServer {
	return &GrpcServer{
		host: host,
		port: port,
	}
}
