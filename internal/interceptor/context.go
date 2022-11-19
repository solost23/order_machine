package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

const request_body = "REQUEST_BODY"

func Context() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, err
	}
}
