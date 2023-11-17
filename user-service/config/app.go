package config

import (
	"log"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func InitApp(conf *AppConfig) {
	env := os.Getenv("GO_ENV")

	switch cases.Lower(language.English).String(env) {
	case "development":
		conf.App.Env = "development"
		log.Println("application running on development environment")
	case "production":
		conf.App.Env = "production"
		log.Println("application running on production environment")
	default:
		conf.App.Env = "development"
		log.Println("application running on development environment")
	}
	conf.App.Env = env
}
