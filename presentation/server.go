package presentation

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
	"google.golang.org/grpc"
)

func NewServer(logger *logrus.Logger,
	recoverHandlerInterceptor *interceptors.RecoverInterceptor,
	errorHandlerInterceptor *interceptors.ErrorHandlerInterceptor,
	localeInterceptor *interceptors.LocaleInterceptor,
	catanServer catan.CatanServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			recoverHandlerInterceptor.ServerUnaryInterceptor(),
			errorHandlerInterceptor.ServerUnaryInterceptor(),
			localeInterceptor.ServerUnaryInterceptor(),
		),
	)

	catan.RegisterCatanServer(server, catanServer)

	return server
}
