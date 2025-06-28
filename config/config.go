package config

import (
	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig() *Config {
	config := &Config{}
	reader := credentials.NewConfigReader()
	if err := reader.Read(gin.Mode(), config); err != nil {
		panic("Can not load config: " + err.Error())
	}
	return config
}
