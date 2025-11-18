package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbUrl     string `mapstructure:"DB_URL"`
	AppAdress string `mapstructure:"APP_ADDRESS"`
}

func ReadConfig(path string) (config *Config, err error) {
	viper.SetConfigName("app")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	fmt.Println(config)
	return config, nil

}
