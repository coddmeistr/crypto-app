package config

import (
	"github.com/spf13/viper"
)

var config *viper.Viper

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

	config = v
	return nil
}

func GetConfig() *viper.Viper {
	return config
}
