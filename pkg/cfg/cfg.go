package cfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
	AWS           AWS           `yaml:"aws"`
}

type Elasticsearch struct {
	Endpoint string `yaml:"endpoint"`
}

type AWS struct {
	Region  string `yaml:"region"`
	RoleARN string `yaml:"roleARN"`
}

func ParseConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error reading config file, %s", err))
	}
	var config *Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error unmarshalling config file, %s", err))
	}
	return config, nil
}
