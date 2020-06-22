package user

import (
	"encoding/json"
	"fmt"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/http"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/util/globals"
)

const userAPIEndpoint = "_opendistro/_security/api/internalusers"

var config = globals.GetConfig()

// Create will create all User using the http package
func Create() error {
	for _, user := range config.Users {
		for key, value := range user {
			jsonUser, err := json.Marshal(value)
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}

			url := fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, userAPIEndpoint, key)
			body := string(jsonUser)

			err = http.FulfillRequest(url, body, "es")
			if err != nil {
				return fmt.Errorf("for user: %s - %s", key, err)
			}
		}
	}
	return nil
}
