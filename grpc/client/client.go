package client

import (
	"projects/Repairment_service/Repairment_service_api_gateway/config"
	"projects/Repairment_service/Repairment_service_api_gateway/genproto/order_service"
	"projects/Repairment_service/Repairment_service_api_gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
	OrderService() order_service.OrderServiceClient
}

type grpcClients struct {
	userService  user_service.UserServiceClient
	orderService order_service.OrderServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:  user_service.NewUserServiceClient(connUserService),
		orderService: order_service.NewOrderServiceClient(connOrderService),
	}, nil
}

func (c *grpcClients) UserService() user_service.UserServiceClient {
	return c.userService
}

func (c *grpcClients) OrderService() order_service.OrderServiceClient {
	return c.orderService
}
