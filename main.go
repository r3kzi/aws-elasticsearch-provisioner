package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

var config *Config
var creds *credentials.Credentials

const (
	service         = "es"
	rolesURL        = "_opendistro/_security/api/roles"
	usersURL        = "_opendistro/_security/api/internalusers"
	rolesMappingURL = "_opendistro/_security/api/rolesmapping"
)

func main() {
	config, _ = parseConfig("config.yml")
	creds = stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)
	if err := createUser(); err != nil {
		fmt.Print(err)
	}
}
