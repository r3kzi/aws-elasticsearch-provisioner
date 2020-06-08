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
	config, _ := cfg.ParseConfig("config.yml")
	creds := stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)

	//Creating user
	users := user.Must(user.Read("./files/users.yml"))
	if err := user.Create(users, config, creds); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
