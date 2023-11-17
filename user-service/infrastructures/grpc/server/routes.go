package server

import (
	"github.com/FianGumilar/microservices/user-service/infrastructures/grpc/rpc/pb"
)

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	pb.RegisterUserServiceServer(rpc.grpcServer, pb.UnimplementedUserServiceServer{})
}
