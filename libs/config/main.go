package config

import "os"

type ACMarketplaceApiConfig struct {
	Host   string
	ApiKey string
}

type InternalConfig struct {
	Environment string
}

type Config struct {
	InternalConfig         InternalConfig
	ACMarketplaceApiConfig ACMarketplaceApiConfig
}

func NewConfig() *Config {
	return &Config{
		InternalConfig: InternalConfig{
			Environment: os.Getenv("ENVIRONMENT"),
		},
		ACMarketplaceApiConfig: ACMarketplaceApiConfig{
			Host:   os.Getenv("AC_MARKETPLACE_API_HOST"),
			ApiKey: os.Getenv("AC_MARKETPLACE_API_KEY"),
		},
	}
}
