package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/r3kzi/elasticsearch-provisioner/cfg"
	"net/http"
	"time"
)

var config *cfg.Config
var creds *credentials.Credentials

const (
	service         = "es"
	rolesURL        = "_opendistro/_security/api/roles"
	usersURL        = "_opendistro/_security/api/internalusers"
	rolesMappingURL = "_opendistro/_security/api/rolesmapping"
)

func main() {
	config, _ = cfg.ParseConfig("config.yml")
	creds = stscreds.NewCredentials(session.Must(session.NewSession()), config.AWS.RoleARN)
	if err := createUser(); err != nil {
		fmt.Print(err)
	}
}

func do(url string, body *bytes.Reader, creds *credentials.Credentials) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating request, %s", err))
	}

	req.Header.Add("Content-Type", "application/json")

	signer := v4.NewSigner(creds)
	_, err = signer.Sign(req, body, service, config.AWS.Region, time.Now())
	if err != nil {
		return errors.New(fmt.Sprintf("Error signing request, %s", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Error making api call, %s", err))
	}
	fmt.Print(resp.Status + "\n")
	return nil
}
