package cfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// Config is the root config for AWS Elasticsearch Service Provisioner
type Config struct {
	Elasticsearch Elasticsearch     `yaml:"elasticsearch"`
	AWS           AWS               `yaml:"aws"`
	Users         []map[string]User `yaml:"users,flow,omitempty"`
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

// User correspond to a Kibana User
type User struct {
	Password     string   `yaml:"password" json:"password"`
	BackendRoles []string `yaml:"backend_roles" json:"backend_roles"`
}

// ParseConfig will parse a configuration file
func ParseConfig(filename string) (*Config, error) {
	ext := filepath.Ext(filename)
	if ext != ".yml" {
		return nil, fmt.Errorf("extension was %s - only yml is supported", ext)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while reading config file, %s", err)
	}

	var config *Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("error while unmarshalling config file, %s", err)
	}

	return config, nil
}
