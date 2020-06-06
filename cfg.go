package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func parseConfig(filename string) (*Config, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading config file, %s", err))
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to decode into struct, %v", err))
	}
	return config, nil
}
