package config

import "github.com/spf13/viper"

type Config struct {
	viper.Viper
}

func NewConfig() *Config {

	// Instancia a struct Vyper
	config := viper.New()

	config.AddConfigPath("./config")
	config.SetConfigType("yml")
	config.SetConfigName("local.yml") // TODO: Change by env
	config.SetEnvPrefix("PREFIX")
	config.ReadInConfig()
	config.AutomaticEnv()
	config.WatchConfig()

	return &Config{*config}
}
