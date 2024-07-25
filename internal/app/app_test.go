package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
	"zg_generator/internal/app/generator"
	"zg_generator/internal/app/grpc_client"
	"zg_generator/internal/app/log"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
		fx.Options(
			generator.NewModule(),
			grpc_client.NewModule(),
			log.NewModule(),
		),
		fx.Provide(
			NewConfig,
		),
	)
	require.NoError(t, err)
}
