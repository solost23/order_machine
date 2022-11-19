package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("gRPC method: %s, %v", info.FullMethod, req)
		resp, err = handler(ctx, req)
		log.Println("resp is:", resp)
		return resp, err
	}
}
