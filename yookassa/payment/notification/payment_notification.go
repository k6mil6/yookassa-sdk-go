package notification

import (
	yoocommon "github.com/k6mil6/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/k6mil6/yookassa-sdk-go/yookassa/payment"
)

type PaymentNotification struct {
	Type   Type               `json:"type"`
	Event  yoocommon.Event    `json:"event"`
	Object yoopayment.Payment `json:"object"`
}
