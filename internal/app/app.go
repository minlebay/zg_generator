package app

import (
	"go.uber.org/fx"
	"zg_generator/internal/app/generator"
	"zg_generator/internal/app/grpc_client"
	"zg_generator/internal/app/log"
	"zg_generator/internal/app/telemetry"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			generator.NewModule(),
			grpc_client.NewModule(),
			log.NewModule(),
			telemetry.NewModule(),
		),
		fx.Provide(
			NewConfig,
		),
	)
}
