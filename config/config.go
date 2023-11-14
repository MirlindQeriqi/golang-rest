package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const ConfFile = "app"
const ConfType = "yml"

func LoadConfig() {
	viper.SetConfigName(ConfFile)
	viper.SetConfigType(ConfType)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetStringConfig(variableName string) string {
	value, success := viper.Get(variableName).(string)

	if !success {
		fmt.Println("Could not read variable: " + variableName)
	}

	return value
}

func GetIntConfig(variableName string) int {
	value, success := viper.Get(variableName).(int)

	if !success {
		fmt.Println("Could not read variable: " + variableName)
	}

	return value
}
