package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"testing"
	"zg_generator/internal/app/generator"
	"zg_generator/internal/app/grpc_client"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
		fx.Options(
			generator.NewModule(),
			grpc_client.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
	require.NoError(t, err)
}
