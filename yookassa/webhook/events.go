package webhook

type Event string

const (
	PaymentWaitingForCapture Event = "payment.waiting_for_capture"
	PaymentSucceeded         Event = "payment.succeeded"
	PaymentCanceled          Event = "payment.canceled"
	RefundSucceeded          Event = "refund.succeeded"
	PayoutSucceeded          Event = "payout.succeeded"
	PayoutCanceled           Event = "payout.canceled"
	DealClosed               Event = "deal.closed"
)
