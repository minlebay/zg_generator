package grpc_client

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	return fx.Module(
		"client",
		fx.Provide(
			NewClientConfig,
			NewClient,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, g *Client) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						go g.StartClient()
						return nil
					},
					OnStop: func(context.Context) error {
						g.StopClient()
						return nil
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("client")
		}),
	)
}
