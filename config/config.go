package config

import "github.com/spf13/viper"

// Config stores all configuration of the application
// THe values are read by viper from a config file or environment variables
type Config struct{
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variable
func LoadConfig(path string, mode string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName(mode)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return
}