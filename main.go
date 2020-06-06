package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/http"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/user"
)

const (
	service         = "es"
	rolesURL        = "_opendistro/_security/api/roles"
	usersURL        = "_opendistro/_security/api/internalusers"
	rolesMappingURL = "_opendistro/_security/api/rolesmapping"
)

func main() {
	config, _ := cfg.ParseConfig("config.yml")
	creds := stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)

	users, _ := user.ReadUser()
	for key, value := range users {
		jsonUser, _ := json.Marshal(value)
		url := fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, usersURL, key)
		body := string(jsonUser)
		request, _ := http.NewRequest(url, string(jsonUser))
		signRequest, _ := http.SignRequest(request, body, creds, service, config.AWS.Region)
		response, _ := http.DoRequest(signRequest)
		fmt.Print(response.Status)
	}
}
