package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env string
	}
	Grpc struct {
		Port string
		Host string
	}
	Postgres struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	if appConfig == nil {
		appConfig = &AppConfig{}
		initPostgres(appConfig)
		InitApp(appConfig)
		initGrpc(appConfig)
	}

	return appConfig
}
