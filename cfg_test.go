package main

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
`

func TestParseConfig(t *testing.T) {
	filename := "test-config.yml"

	bytes := []byte(data)
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		t.Errorf("Failed to write file, %v", err)
	}

	config, err := parseConfig(filename)
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, config.Elasticsearch.Endpoint, "https://elasticsearch:443")
	assert.Equal(t, config.AWS.Region, "eu-west-1")
	assert.Equal(t, config.AWS.RoleARN, "arn:aws:iam::123456789012:role/IAMMasterUser")

	if err := os.Remove(filename); err != nil {
		t.Errorf("Failed to remove file, %v", err)
	}
}
