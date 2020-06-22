package globals

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
)

var config *cfg.Config = &cfg.Config{}
var cred *credentials.Credentials = &credentials.Credentials{}

func GetConfig() *cfg.Config {
	return config
}

func GetCredentials() *credentials.Credentials {
	return cred
}

func SetConfig(configuration *cfg.Config) {
	config = configuration
}

func SetCredentials(credentials *credentials.Credentials) {
	cred = credentials
}
