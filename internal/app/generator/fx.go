package generator

import (
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
				lc.Append(fx.StartStopHook(g.StartGenerator, g.StopGenerator))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("generator")
		}),
	)
}
