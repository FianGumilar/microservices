package server

import "github.com/FianGumilar/microservices/user-service/infrastructures/grpc/handler"

type gRPCProvider struct {
	handlers struct {
		user handler.UserHandler
	}
}
