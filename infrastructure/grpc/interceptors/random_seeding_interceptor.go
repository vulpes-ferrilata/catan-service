package interceptors

import (
	"context"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

func NewRandomSeedingInterceptor() *RandomSeedingInterceptor {
	return &RandomSeedingInterceptor{}
}

type RandomSeedingInterceptor struct{}

func (r RandomSeedingInterceptor) ServerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		rand.Seed(time.Now().UnixNano())

		return handler(ctx, req)
	}
}
