package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}

	Environment Environment `mapstructure:"env"`

	CryptoCompare CryptoCompare `mapstructure:"crypto_compare"`
}

type Environment struct {
	Mode string
}

type CryptoCompare struct {
	AppName string `mapstructure:"app_name"`
}

var config Config

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) error {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)

	v.AddConfigPath("../config/")
	v.AddConfigPath("config/")
	v.AddConfigPath("crypto-server/config/")
	v.AddConfigPath("crypto/config/")

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	config = cfg
	return nil
}

func GetConfig() Config {
	return config
}
