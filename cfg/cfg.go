package cfg

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Elasticsearch Elasticsearch
	AWS           AWS
}

type Elasticsearch struct {
	Endpoint string
}

type AWS struct {
	Region  string
	RoleARN string
}

func ParseConfig(filename string) (*Config, error) {
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
