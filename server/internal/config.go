package internal

import "github.com/spf13/viper"

type Config struct {
	Address     string `mapstructure:"ADDRESS"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}