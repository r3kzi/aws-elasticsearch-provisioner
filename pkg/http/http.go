package http

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"strings"
	"time"
)

func NewRequest(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating request, %s", err))
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func SignRequest(req *http.Request, body string, creds *credentials.Credentials, service string, region string) (*http.Request, error) {
	signer := v4.NewSigner(creds)
	_, err := signer.Sign(req, strings.NewReader(body), service, region, time.Now())
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error signing request, %s", err))
	}
	return req, nil
}

func DoRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error making api call, %s", err))
	}
	return resp, nil
}
