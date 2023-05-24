package agoraio

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.agora.io/dev"

type client struct {
	baseURL        string
	appID          string
	customerKey    string
	customerSecret string
	httpClient     *http.Client
}

func NewClient(appID, customerKey, customerSecret string) *client {
	return &client{
		baseURL:        baseURL,
		appID:          appID,
		customerKey:    customerKey,
		customerSecret: customerSecret,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) request(method, endpoint string, params, response interface{}) error {
	var body io.Reader
	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(data)
	}

	request, err := http.NewRequest(method, baseURL+endpoint, body)
	if err != nil {
		return err
	}

	plainCredentials := c.customerKey + ":" + c.customerSecret
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(plainCredentials))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+base64Credentials)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
