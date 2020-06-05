package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"
)

var config *Config

const (
	service         = "es"
	rolesURL        = "_opendistro/_security/api/roles"
	usersURL        = "_opendistro/_security/api/internalusers"
	rolesMappingURL = "_opendistro/_security/api/rolesmapping"
)

func readConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return config
}

func readUser() map[string]User {
	file, err := ioutil.ReadFile("./files/users.yml")
	if err != nil {
		fmt.Printf("Error reading users file, %s", err)
	}

	var users map[string]User
	if err := yaml.Unmarshal(file, &users); err != nil {
		fmt.Printf("Error unmarshaling users file, %s", err)
	}
	return users
}

func do(url string, body *bytes.Reader, creds *credentials.Credentials) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		fmt.Printf("Error creating request, %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	signer := v4.NewSigner(creds)
	_, err = signer.Sign(req, body, service, config.AWS.Region, time.Now())
	if err != nil {
		fmt.Printf("Error signing request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making api call, %s", err)
	}
	fmt.Print(resp.Status + "\n")
}

func main() {
	config = readConfig()

	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, config.AWS.RoleARN)

	for key, user := range readUser() {
		jsonUser, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("Error marshalling json for %s, %s", key, err)
		}
		do(fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, usersURL, key), bytes.NewReader(jsonUser), creds)
	}
}
