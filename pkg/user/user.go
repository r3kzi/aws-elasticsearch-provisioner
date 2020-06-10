package user

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/http"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const userAPIEndpoint = "_opendistro/_security/api/internalusers"

// User correspond a Kibana User
type User struct {
	Password     string   `yaml:"password" json:"password"`
	BackendRoles []string `yaml:"backend_roles" json:"backend_roles"`
}

// Read will a file und tries to parse User structs
func Read(filename string) (map[string]User, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error reading users file, %s", err))
	}

	var users map[string]User
	if err := yaml.Unmarshal(file, &users); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error unmarshalling users file, %s", err))
	}
	return users, nil
}

// Create will create all User using the http package
func Create(users map[string]User, config *cfg.Config, creds *credentials.Credentials) error {
	for key, value := range users {
		jsonUser, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("For user: %s, %s", key, err))
		}

		url := fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, userAPIEndpoint, key)
		body := string(jsonUser)

		request, err := http.NewRequest(url, string(jsonUser))
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("For user: %s, %s", key, err))
		}

		signRequest, err := http.SignRequest(request, body, creds, "es", config.AWS.Region)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("For user: %s, %s", key, err))
		}

		resp, err := http.DoRequest(signRequest)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("For user: %s, %s", key, err))
		}

		if resp.StatusCode != 200 && resp.StatusCode != 201 {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			return fmt.Errorf(fmt.Sprintf("Failed creating user: %s, Status Code was %d\n"+
				"Message: %s", key, resp.StatusCode, bodyString))
		}
	}
	return nil
}

// Must ensures validate user map
func Must(users map[string]User, err error) map[string]User {
	if err != nil {
		panic(err)
	}
	return users
}
