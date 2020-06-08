package http

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	url     = "https://elasticsearch"
	body    = "body"
	service = "es"
	region  = "eu-west-1"
)

type CredentialsTestProvider struct{}

func (c *CredentialsTestProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{}, nil
}
func (c *CredentialsTestProvider) IsExpired() bool {
	return false
}

func TestNewRequest(t *testing.T) {
	request, err := NewRequest(url, body)
	assert.Nil(t, err)
	assert.Equal(t, request.URL.Host, "elasticsearch")
	assert.Equal(t, request.URL.Scheme, "https")

	bytes, err := ioutil.ReadAll(request.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), "body")
}

func TestSignRequest(t *testing.T) {
	creds := credentials.NewCredentials(&CredentialsTestProvider{})
	request, err := NewRequest(url, body)
	assert.Nil(t, err)

	signRequest, err := SignRequest(request, body, creds, service, region)
	assert.Nil(t, err)
	assert.NotNil(t, signRequest)
	assert.NotEmpty(t, signRequest.Header.Get("Authorization"))
}
