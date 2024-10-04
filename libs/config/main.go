package config

import "os"

type ACMarketplaceApiConfig struct {
	Host   string
	ApiKey string
}

type Auth0 struct {
	JwtSecret string
}

type InternalConfig struct {
	Environment string
}

type Config struct {
	ACMarketplaceApiConfig ACMarketplaceApiConfig
	Auth0                  Auth0
	InternalConfig         InternalConfig
}

func NewConfig() *Config {
	return &Config{
		ACMarketplaceApiConfig: ACMarketplaceApiConfig{
			Host:   os.Getenv("AC_MARKETPLACE_API_HOST"),
			ApiKey: os.Getenv("AC_MARKETPLACE_API_KEY"),
		},
		Auth0: Auth0{
			JwtSecret: os.Getenv("AUTH0_JWT_SECRET"),
		},
		InternalConfig: InternalConfig{
			Environment: os.Getenv("ENVIRONMENT"),
		},
	}
}
