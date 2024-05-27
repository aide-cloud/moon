package server

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/aide-cloud/moon/cmd/server/rabbit/internal/rabbitconf"
	"github.com/aide-cloud/moon/pkg/helper/middleware"
	"github.com/aide-cloud/moon/pkg/log"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bc *rabbitconf.Bootstrap) *grpc.Server {
	c := bc.GetServer()
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			middleware.Logging(log.GetLogger()),
			middleware.Validate(protovalidate.WithFailFast(true)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	return srv
}