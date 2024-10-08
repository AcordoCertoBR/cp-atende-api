package config

import "os"

type ACMarketplaceApiConfig struct {
	Host   string
	ApiKey string
}

type Auth0 struct {
	PublicCertificate string
}

type InternalConfig struct {
	Environment string
	LogLevel    string
	AppName     string
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
			PublicCertificate: os.Getenv("AUTH0_PUBLIC_CERTIFICATE"),
		},
		InternalConfig: InternalConfig{
			Environment: os.Getenv("ENVIRONMENT"),
			LogLevel:    os.Getenv("LOG_LEVEL"),
			AppName:     "cp-atende-api",
		},
	}
}
