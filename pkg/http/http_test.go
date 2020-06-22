package http

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/util/globals"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

const (
	url     = "https://elasticsearch.com"
	body    = "body"
	service = "es"
)

var configuration = &cfg.Config{
	Elasticsearch: cfg.Elasticsearch{},
	AWS:           cfg.AWS{Region: "eu-west-1"},
	Users:         []map[string]cfg.User{},
}

func init() {
	globals.SetConfig(configuration)
	globals.SetCredentials(credentials.NewCredentials(&CredentialsTestProvider{}))
}

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
	request, err := NewRequest(url, body)
	assert.Nil(t, err)

	signRequest, err := SignRequest(request, service)
	assert.Nil(t, err)
	assert.NotNil(t, signRequest)
	assert.NotEmpty(t, signRequest.Header.Get("Authorization"))
	assert.Regexp(t, regexp.MustCompile("^AWS4-HMAC-SHA256"), signRequest.Header.Get("Authorization"))
}

func TestDoRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	req, _ := http.NewRequest(http.MethodPut, server.URL, strings.NewReader(""))
	err := DoRequest(req)
	assert.Nil(t, err)
}

func TestFulfillRequest(t *testing.T) {
	// TODO: not really required while FulfillRequest method is just a composition of the methods above
	/*
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		err := FulfillRequest(server.URL, "", service)
		assert.Nil(t, err)
	*/
}
