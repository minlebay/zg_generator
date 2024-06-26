package generator

import (
	"go.uber.org/config"
)

type Config struct {
	Interval      int    `yaml:"interval"`
	Count         int    `yaml:"count"`
	RouterAddress string `yaml:"router_address"`
}

func NewGeneratorConfig(provider config.Provider) (*Config, error) {
	var config Config
	err := provider.Get("generator").Populate(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
