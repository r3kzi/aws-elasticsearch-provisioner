package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/user"
	"os"
)

func main() {
	configPath := flag.String("config-path", "config.yml", "Path of the configuration file - has to be yml format")
	flag.Parse()

	config, err := cfg.ParseConfig(*configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// NewCredentials returns a pointer to a new Credentials object wrapping the AssumeRoleProvider
	creds := stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)

	// Creating user
	if err := user.Create(config, creds); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creating roles
	// Creating rolesmapping
}
