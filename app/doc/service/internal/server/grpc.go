package server

import (
	"crypto/tls"

	conf "github.com/ToAtlas/AtlasBackend/api/gen/go/conf/v1"
	docv1 "github.com/ToAtlas/AtlasBackend/api/gen/go/doc/service/v1"
	"github.com/ToAtlas/AtlasBackend/app/doc/service/internal/service"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, doc *service.DocService) *grpc.Server {
	var mds []middleware.Middleware
	mds = []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(logger),
		validate.ProtoValidate(),
	}

	var opts = []grpc.ServerOption{
		grpc.Middleware(mds...),
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

	if c.Grpc.Tls != nil && c.Grpc.Tls.Enable {
		cert, err := tls.LoadX509KeyPair(c.Grpc.Tls.CertPath, c.Grpc.Tls.KeyPath)
		if err != nil {
			logger.Log(log.LevelFatal, "msg", "gRPC Server TLS: Failed to load key pair", "error", err)
		}
		creds := credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}})
		opts = append(opts, grpc.Options(gogrpc.Creds(creds)))
	}

	srv := grpc.NewServer(opts...)
	// 注意：以生成代码中的注册函数名为准
	docv1.RegisterDocServer(srv, doc)
	return srv
}
