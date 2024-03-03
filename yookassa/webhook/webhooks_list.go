package webhook

type WebhooksList struct {
	Type     string    `json:"type"`
	Webhooks []Webhook `json:"webhooks"`
}
