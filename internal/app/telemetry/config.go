package telemetry

import cfg "go.uber.org/config"

type Config struct {
	Url string `yaml:"url"`
}

func NewMetricsConfig(provider cfg.Provider) (*Config, error) {
	var config Config
	err := provider.Get("prometheus").Populate(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
