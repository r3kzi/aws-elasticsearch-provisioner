package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"time"
)

func do(url string, body *bytes.Reader, creds *credentials.Credentials) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating request, %s", err))
	}

	req.Header.Add("Content-Type", "application/json")

	signer := v4.NewSigner(creds)
	_, err = signer.Sign(req, body, service, config.AWS.Region, time.Now())
	if err != nil {
		return errors.New(fmt.Sprintf("Error signing request, %s", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Error making api call, %s", err))
	}
	fmt.Print(resp.Status + "\n")
	return nil
}
