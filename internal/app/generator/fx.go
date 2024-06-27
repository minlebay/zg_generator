package generator

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	return fx.Module(
		"generator",
		fx.Provide(
			NewGeneratorConfig,
			NewGenerator,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, g *Generator) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go g.StartGenerator(ctx)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						g.StopGenerator(ctx)
						return nil
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("generator")
		}),
	)
}
