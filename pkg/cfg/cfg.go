package cfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config is the root config for AWS Elasticsearch Service Provisioner
type Config struct {
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
	AWS           AWS           `yaml:"aws"`
}

// Elasticsearch contains all necessary configuration for an Elasticsearch Domain
type Elasticsearch struct {
	Endpoint string `yaml:"endpoint"`
}

// AWS contains all AWS-specific configuration
type AWS struct {
	Region  string `yaml:"region"`
	RoleARN string `yaml:"roleARN"`
}

// ParseConfig will parse a configuration file
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
