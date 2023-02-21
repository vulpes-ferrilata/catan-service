package presentation

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/vulpes-ferrilata/catan-service-proto/pb"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/grpc/interceptors"
	"google.golang.org/grpc"
)

func NewServer(logger *logrus.Logger,
	recoverHandlerInterceptor *interceptors.RecoverInterceptor,
	errorHandlerInterceptor *interceptors.ErrorHandlerInterceptor,
	localeInterceptor *interceptors.LocaleInterceptor,
	randomSeedingInterceptor *interceptors.RandomSeedingInterceptor,
	catanServer pb.CatanServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			recoverHandlerInterceptor.ServerUnaryInterceptor(),
			randomSeedingInterceptor.ServerUnaryInterceptor(),
			localeInterceptor.ServerUnaryInterceptor(),
			errorHandlerInterceptor.ServerUnaryInterceptor(),
		),
	)

	pb.RegisterCatanServer(server, catanServer)

	return server
}
