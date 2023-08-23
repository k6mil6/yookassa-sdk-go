package yookassa

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const (
	BaseURL = "https://api.yookassa.ru/v3/"
)

type Client struct {
	client    http.Client
	accountId string
	secretKey string
}

func NewClient(accountId string, secretKey string) *Client {
	return &Client{
		client:    http.Client{},
		accountId: accountId,
		secretKey: secretKey,
	}
}

func (c *Client) makeRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", BaseURL, endpoint)

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotence-Key", uuid.NewString())
	}

	req.SetBasicAuth(c.accountId, c.secretKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}