package grpc_client

import "go.uber.org/config"

type Config struct {
	RouterAddress string `yaml:"router_address"`
}

func NewClientConfig(provider config.Provider) (*Config, error) {
	var config Config
	err := provider.Get("grpc_client").Populate(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
