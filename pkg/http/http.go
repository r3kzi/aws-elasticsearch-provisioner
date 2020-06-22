package http

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/r3kzi/elasticsearch-provisioner/pkg/util/globals"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var region = globals.GetConfig().AWS.Region
var signer = v4.NewSigner(globals.GetCredentials())

// NewRequest will create a new HTTP request based on an URL and a Body string
func NewRequest(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating http request, %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// SignRequest will sign a HTTP requests with an assumed role for a specific AWS Region
// using AWS Signature V4 signing process
func SignRequest(req *http.Request, service string) (*http.Request, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	_, err = signer.Sign(req, bytes.NewReader(body), service, region, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error signing http request, %s", err)
	}
	return req, nil
}

// DoRequest will actually make the http call
func DoRequest(req *http.Request) error {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making elasticsearch api call: %s", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("returned code was %s - message was %s", resp.Status, string(bytes))
	}
	return nil
}

// FulfillRequest is a helper function which covers all required steps
func FulfillRequest(url string, body string, service string) error {
	request, err := NewRequest(url, body)
	if err != nil {
		return err
	}

	signRequest, err := SignRequest(request, service)
	if err != nil {
		return err
	}

	err = DoRequest(signRequest)
	if err != nil {
		return err
	}

	return nil
}
