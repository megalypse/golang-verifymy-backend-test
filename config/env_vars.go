package config

import "os"

var secret string

func init() {
	secret = os.Getenv("AUTH_SECRET")
}

func GetAuthSecret() string {
	return secret
}
