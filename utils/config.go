package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE"`
	TokenSecretKey      string        `mapstructure:"TOKEN_SECRET_KEY"`
	GRPCServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	HTTPCServerAddress  string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	WebhookVerifyToken  string        `mapstructure:"WEBHOOK_VERIFY_TOKEN"`
	MetaApiToken        string        `mapstructure:"META_API_TOKEN"`
	MetaPhoneNumberID   string        `mapstructure:"META_PHONE_NUMBER_ID"`
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
