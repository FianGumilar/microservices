package config

import "os"

func initGrpc(conf *AppConfig) {
	port := os.Getenv("GRPC_PORT")
	host := os.Getenv("GRPC_HOST")

	conf.Grpc.Port = port
	conf.Grpc.Host = host
}
