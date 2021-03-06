package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config stores all configurations for the application
type Config struct {
	DBDriver      string `mapstructure:"db_driver"`
	DBSource      string `mapstructure:"db_source"`
	ServerAddress string `mapstructure:"server_address"`
}

// LoadConfig reads configurations from a file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetDefault("db_driver", "postgres")
	viper.SetDefault("db_source", "user=root password=secret host=localhost dbname=gobutsu sslmode=disable")
	viper.SetDefault("server_address", "0.0.0.0:8081")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Using DEFAULT values.", err)
	}

	err = viper.Unmarshal(&config)
	return
}
