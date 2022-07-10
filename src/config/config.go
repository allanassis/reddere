package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	viper.Viper
}

func NewConfig() *Config {
	env, hasEnv := os.LookupEnv("ENV")
	if !hasEnv {
		panic("Please set the ENV variable")
	}
	config := viper.New()

	config.AddConfigPath("./config")
	config.SetConfigType("yml")
	config.SetConfigName(fmt.Sprintf("%s.yml", env))
	config.SetEnvPrefix("PREFIX")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config.AutomaticEnv()
	config.WatchConfig()

	return &Config{*config}
}
