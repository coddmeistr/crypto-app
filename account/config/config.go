package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}

	Environment Environment `mapstructure:"env"`

	Validation struct {
		Password struct {
			MinLength int `mapstructure:"min_length"`
			MaxLength int `mapstructure:"max_length"`
		}
		Login struct {
			MinLength int `mapstructure:"min_length"`
			MaxLength int `mapstructure:"max_length"`
		}
	}

	Dependencies struct {
		CryptoService CryptoService `mapstructure:"crypto_service"`
	}
}

type Environment struct {
	Mode string
}

type CryptoService struct {
	Host      string
	Endpoints struct {
		GetCurrentPrices string `mapstructure:"current_prices"`
	}
}

func getCryptoServiceHost() (string, error) {
	host := os.Getenv("CRYPTO_SERVICE_HOST")
	if host == "" {
		return host, errors.New("CRYPTO_SERVICE_HOST env is empty")
	}
	return host, nil
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
	v.AddConfigPath("account-app/config/")
	v.AddConfigPath("account-app/account/config/")

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	cryptoHost, err := getCryptoServiceHost()
	if err != nil {
		return err
	}
	cfg.Dependencies.CryptoService.Host = cryptoHost

	config = cfg
	return nil
}

func GetConfig() Config {
	return config
}
