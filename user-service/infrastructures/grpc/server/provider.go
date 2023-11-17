package server

import (
	"github.com/FianGumilar/microservices/user-service/config"
	"github.com/FianGumilar/microservices/user-service/infrastructures/database"
	"github.com/FianGumilar/microservices/user-service/infrastructures/grpc/handler"
	"github.com/FianGumilar/microservices/user-service/repositories"
	"github.com/FianGumilar/microservices/user-service/services"
)

type gRPCProvider struct {
	handlers struct {
		user handler.UserHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	conf := config.NewAppConfig()
	db := database.NewPostgres(conf)

	userRepo := repositories.NewUserRespository(db)
	userService := services.NewUserService(userRepo)
	provider.handlers.user = *handler.NewUserHandlerUserServices(userService)

	return provider
}
