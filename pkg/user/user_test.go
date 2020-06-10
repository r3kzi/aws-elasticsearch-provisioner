package user

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var config = cfg.Config{
	Elasticsearch: cfg.Elasticsearch{},
	AWS:           cfg.AWS{},
	Users: []map[string]cfg.User{
		{
			"max":   cfg.User{Password: "password", BackendRoles: []string{"backend-role-1"}},
			"erica": cfg.User{Password: "password", BackendRoles: []string{"backend-role-2", "backend-role-3"}},
		},
	},
}

type CredentialsTestProvider struct{}

func (c *CredentialsTestProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{}, nil
}
func (c *CredentialsTestProvider) IsExpired() bool {
	return false
}

func TestCreateStatusOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Regexp(t, regexp.MustCompile("/_opendistro/_security/api/internalusers/*"), r.URL.Path)
		assert.Regexp(t, regexp.MustCompile("^AWS4-HMAC-SHA256"), r.Header["Authorization"][0])

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}))
	defer server.Close()

	config.Elasticsearch.Endpoint = server.URL

	creds := credentials.NewCredentials(&CredentialsTestProvider{})
	err := Create(&config, creds)
	assert.Nil(t, err)
}

func TestCreateStatusBadRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Regexp(t, regexp.MustCompile("/_opendistro/_security/api/internalusers/*"), r.URL.Path)
		assert.Regexp(t, regexp.MustCompile("^AWS4-HMAC-SHA256"), r.Header["Authorization"][0])

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}))
	defer server.Close()

	config.Elasticsearch.Endpoint = server.URL

	creds := credentials.NewCredentials(&CredentialsTestProvider{})
	err := Create(&config, creds)
	assert.NotNil(t, err)
}
