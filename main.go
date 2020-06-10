package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/user"
	"os"
)

func main() {
	config, err := cfg.ParseConfig("config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// NewCredentials returns a pointer to a new Credentials object wrapping the AssumeRoleProvider
	creds := stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)

	// Creating user
	users := user.Must(user.Read("./files/users.yml"))
	if err := user.Create(users, config, creds); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
