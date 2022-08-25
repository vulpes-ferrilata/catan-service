package presentation

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
	"google.golang.org/grpc"
)

func NewServer(logger *logrus.Logger,
	errorHandlerInterceptor *interceptors.ErrorHandlerInterceptor,
	localeInterceptor *interceptors.LocaleInterceptor,
	catanServer catan.CatanServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			errorHandlerInterceptor.ServerUnaryInterceptor,
			localeInterceptor.ServerUnaryInterceptor,
		),
	)

	catan.RegisterCatanServer(server, catanServer)

	return server
}
