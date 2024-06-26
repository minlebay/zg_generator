package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_generator/internal/app/generator"
	"zg_generator/internal/app/grpc_client"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			generator.NewModule(),
			grpc_client.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
}
