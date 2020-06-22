package user

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/cfg"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/util/globals"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var configuration = &cfg.Config{
	Elasticsearch: cfg.Elasticsearch{},
	AWS:           cfg.AWS{},
	Users: []map[string]cfg.User{
		{
			"max":   cfg.User{Password: "password", BackendRoles: []string{"backend-role-1"}},
			"erica": cfg.User{Password: "password", BackendRoles: []string{"backend-role-2", "backend-role-3"}},
		},
	},
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

func TestCreate(t *testing.T) {
	tests := []struct {
		StatusCode int
		Handler    http.HandlerFunc
	}{
		{StatusCode: http.StatusOK, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertRequest(t, r)
			w.WriteHeader(http.StatusOK)
		})},
		{StatusCode: http.StatusCreated, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertRequest(t, r)
			w.WriteHeader(http.StatusCreated)
		})},
		{StatusCode: http.StatusBadRequest, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertRequest(t, r)
			w.WriteHeader(http.StatusBadRequest)
		})},
	}

	for _, test := range tests {
		server := httptest.NewServer(test.Handler)

		configuration.Elasticsearch.Endpoint = server.URL

		err := Create()
		switch test.StatusCode {
		case http.StatusOK:
			assert.Nil(t, err)
		case http.StatusCreated:
			assert.Nil(t, err)
		case http.StatusBadRequest:
			assert.NotNil(t, err)
		}
		server.Close()
	}
}

func assertRequest(t *testing.T, r *http.Request) {
	assert.Equal(t, http.MethodPut, r.Method)
	assert.Regexp(t, regexp.MustCompile("/_opendistro/_security/api/internalusers/*"), r.URL.Path)
	assert.Regexp(t, regexp.MustCompile("^AWS4-HMAC-SHA256"), r.Header["Authorization"][0])
}
