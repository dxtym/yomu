package internal

import "github.com/spf13/viper"

type Config struct {
	Address      string `mapstructure:"ADDRESS"`
	ApiUrl       string `mapstructure:"API_URL"`
	RedisAddr    string `mapstructure:"REDIS_ADDR"`
	BotToken     string `mapstructure:"BOT_TOKEN"`
	PostgresAddr string `mapstructure:"POSTGRES_ADDR"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}
