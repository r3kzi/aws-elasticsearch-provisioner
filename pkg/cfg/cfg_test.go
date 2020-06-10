package cfg

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var data = `
elasticsearch:
  endpoint: "https://elasticsearch:443"

aws:
  region: "eu-west-1"
  roleARN: "arn:aws:iam::123456789012:role/IAMMasterUser"

users:
  - max:
      password: I07%0&8bv4$ie!92hr7q6fxs#7wUGO9%jJqV
      backend_roles:
        - backend-role-1
        - backend-role-2
`

func TestParseConfig(t *testing.T) {
	filename := "test-config.yml"

	bytes := []byte(data)
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		t.Errorf("failed to write file, %v", err)
	}

	config, err := ParseConfig(filename)
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "https://elasticsearch:443", config.Elasticsearch.Endpoint)
	assert.Equal(t, "eu-west-1", config.AWS.Region)
	assert.Equal(t, "arn:aws:iam::123456789012:role/IAMMasterUser", config.AWS.RoleARN)

	assert.NotEmpty(t, config.Users)
	assert.Equal(t, "I07%0&8bv4$ie!92hr7q6fxs#7wUGO9%jJqV", config.Users[0]["max"].Password)
	assert.Equal(t, "backend-role-2", config.Users[0]["max"].BackendRoles[1])

	if err := os.Remove(filename); err != nil {
		t.Errorf("failed to remove file, %v", err)
	}
}

func TestParseConfigWithJson(t *testing.T) {
	filename := "test-config.json"
	_, err := ParseConfig(filename)
	assert.NotNil(t, err)
}
