package yookassa

import (
	"encoding/json"
	"fmt"
	yooerror "github.com/k6mil6/yookassa-sdk-go/yookassa/errors"
	"github.com/k6mil6/yookassa-sdk-go/yookassa/webhook"
	"io"
	"net/http"
)

const (
	WebhookEndpoint = "webhooks"
)

type WebhookHandler struct {
	client *Client
}

func NewWebhookHandler(client *Client) *WebhookHandler {
	return &WebhookHandler{client: client}
}

func (w *WebhookHandler) AddWebhook(webhook *webhook.Webhook) (*webhook.Webhook, error) {
	webhookJson, err := json.Marshal(webhook)
	if err != nil {
		return nil, err
	}

	resp, err := w.client.makeRequestWebhook(http.MethodPost, WebhookEndpoint, webhookJson, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	webhookResponse, err := w.parseWebhookResponse(resp)
	if err != nil {
		return nil, err
	}
	return webhookResponse, nil
}

func (w *WebhookHandler) GetAllWebhooks() ([]webhook.Webhook, error) {
	resp, err := w.client.makeRequestWebhook(http.MethodGet, WebhookEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	webhooks, err := w.parseWebhooksResponse(resp)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (w *WebhookHandler) RemoveWebhook(id string) error {
	endpoint := fmt.Sprintf("%s/%s", WebhookEndpoint, id)
	resp, err := w.client.makeRequestWebhook(http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return err
		}

		return respError
	}

	return nil
}

func (w *WebhookHandler) parseWebhookResponse(resp *http.Response) (*webhook.Webhook, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	webhookResponse := webhook.Webhook{}
	err = json.Unmarshal(responseBytes, &webhookResponse)
	if err != nil {
		return nil, err
	}
	return &webhookResponse, nil
}

func (w *WebhookHandler) parseWebhooksResponse(resp *http.Response) ([]webhook.Webhook, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	webhooksResponse := webhook.WebhooksList{}
	err = json.Unmarshal(responseBytes, &webhooksResponse)
	if err != nil {
		return nil, err
	}
	return webhooksResponse.Webhooks, nil
}
