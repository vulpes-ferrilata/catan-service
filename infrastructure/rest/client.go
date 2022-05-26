package rest

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/v3/client"
)

func NewClient(errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware) client.Client {
	server := grpc.NewClient(
		client.WrapCall(errorHandlerMiddleware.WrapCall),
	)

	return server
}
