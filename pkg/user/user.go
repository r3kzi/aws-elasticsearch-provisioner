package user

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type User struct {
	Password     string   `yaml:"password" json:"password"`
	BackendRoles []string `yaml:"backend_roles" json:"backend_roles"`
}

const filename = "./files/users.yml"

func ReadUser() (map[string]User, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading users file, %s", err))
	}

	var users map[string]User
	if err := yaml.Unmarshal(file, &users); err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling users file, %s", err))
	}
	return users, nil
}
