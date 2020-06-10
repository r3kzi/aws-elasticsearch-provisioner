package http

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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
func SignRequest(req *http.Request, body string, creds *credentials.Credentials, service string, region string) (*http.Request, error) {
	signer := v4.NewSigner(creds)
	_, err := signer.Sign(req, strings.NewReader(body), service, region, time.Now())
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
