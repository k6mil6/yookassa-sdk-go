package webhook

import "github.com/k6mil6/yookassa-sdk-go/yookassa/common"

type Webhook struct {
	ID    string          `json:"id,omitempty"`
	Event yoocommon.Event `json:"event,omitempty"`
	URL   string          `json:"url,omitempty"`
}
