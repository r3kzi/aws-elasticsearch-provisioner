package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const filename = "./files/users.yml"

func createUser() error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading users file, %s", err))
	}

	var users map[string]User
	if err := yaml.Unmarshal(file, &users); err != nil {
		return errors.New(fmt.Sprintf("Error unmarshaling users file, %s", err))
	}

	for key, user := range users {
		jsonUser, err := json.Marshal(user)
		if err != nil {
			return errors.New(fmt.Sprintf("Error marshalling json for %s, %s", key, err))
		}
		return do(fmt.Sprintf("%s/%s/%s", config.Elasticsearch.Endpoint, usersURL, key), bytes.NewReader(jsonUser), creds)
	}
	return nil
}
