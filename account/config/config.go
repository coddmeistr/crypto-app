package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
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
		fmt.Println("panic", err)
	}

	config = v
}

func GetConfig() *viper.Viper {
	return config
}
