package config

import "os"

var secret string
var serverContainerPort string
var serverHostPort string

func init() {
	secret = os.Getenv("AUTH_SECRET")
	serverContainerPort = os.Getenv("SERVER_CONTAINER_PORT")
	serverHostPort = os.Getenv("SERVER_HOST_PORT")
}

func GetAuthSecret() string {
	return secret
}

func GetServerContainerPort() string {
	return serverContainerPort
}

func GetServerHostPort() string {
	return serverHostPort
}
