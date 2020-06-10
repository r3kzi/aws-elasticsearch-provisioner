package user

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/http"
)

const userAPIEndpoint = "_opendistro/_security/api/internalusers"

// Create will create all User using the http package
func Create(config *cfg.Config, creds *credentials.Credentials) error {
	for _, user := range config.Users {
		for key, value := range user {
			jsonUser, err := json.Marshal(value)
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}

			url := fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, userAPIEndpoint, key)
			body := string(jsonUser)

			request, err := http.NewRequest(url, string(jsonUser))
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}

			signRequest, err := http.SignRequest(request, body, creds, "es", config.AWS.Region)
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}

			err = http.DoRequest(signRequest)
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}
		}
	}
	return nil
}
