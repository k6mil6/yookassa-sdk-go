package webhook

type Webhook struct {
	ID    string `json:"id,omitempty"`
	Event Event  `json:"event,omitempty"`
	URL   string `json:"url,omitempty"`
}
