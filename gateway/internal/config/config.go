package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port      string
		SecretKey string
	}

	Auth struct {
		JwtExpSeconds int
	}

	Environment Environment `mapstructure:"env"`

	AccountService AccountService `mapstructure:"account_service"`
	CryptoService  CryptoService  `mapstructure:"crypto_service"`
}

type Environment struct {
	Mode string
}

type AccountService struct {
	Host string
	Url  string
}

type CryptoService struct {
	Host string
	Url  string
	Ws   string
}

func getCryptoServiceHost() (string, error) {
	host := os.Getenv("CRYPTO_SERVICE_HOST")
	if host == "" {
		return host, errors.New("CRYPTO_SERVICE_HOST env is empty")
	}
	return host, nil
}

func getAccountServiceHost() (string, error) {
	host := os.Getenv("ACCOUNT_SERVICE_HOST")
	if host == "" {
		return host, errors.New("ACCOUNT_SERVICE_HOST env is empty")
	}
	return host, nil
}

func getSecretKey() (string, error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return key, errors.New("SECRET_KEY env is empty")
	}
	return key, nil
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
	v.AddConfigPath("../internal/config/")
	v.AddConfigPath("internal/config/")
	v.AddConfigPath("../internal/config/")
	v.AddConfigPath("config/")
	v.AddConfigPath("crypto-app/config/")
	v.AddConfigPath("crypto-app/gateway/config/")

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	// Crypto host
	cryptoHost, err := getCryptoServiceHost()
	if err != nil {
		return err
	}
	cfg.CryptoService.Host = cryptoHost

	// Account host
	accountHost, err := getAccountServiceHost()
	if err != nil {
		return err
	}
	cfg.AccountService.Host = accountHost

	// Secret key
	key, err := getSecretKey()
	if err != nil {
		return err
	}
	cfg.Server.SecretKey = key

	config = cfg
	return nil
}

func GetConfig() Config {
	return config
}
