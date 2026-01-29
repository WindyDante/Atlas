//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ToAtlas/AtlasBackend/api/gen/go/conf/v1"
	"github.com/ToAtlas/AtlasBackend/app/doc/service/internal/server"
	"github.com/ToAtlas/AtlasBackend/app/doc/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
