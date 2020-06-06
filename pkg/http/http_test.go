package http

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	url  = "https://elasticsearch"
	body = "body"
)

func TestNewRequest(t *testing.T) {
	request, err := NewRequest(url, body)
	assert.Nil(t, err)
	assert.Equal(t, request.URL.Host, "elasticsearch")
	assert.Equal(t, request.URL.Scheme, "https")

	bytes, err := ioutil.ReadAll(request.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), "body")
}
