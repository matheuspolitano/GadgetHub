package utils

import "github.com/spf13/viper"

type Config struct {
	DBSource           string `mapstructure:"DB_SOURCE"`
	TokenSecretKey     string `mapstructure:"TOKEN_SECRET_KEY"`
	GRPCServerAddress  string `mapstructure:"GRPC_SERVER_ADDRESS"`
	HTTPCServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig get the configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
