package smsotpverifier

import context "context"

// Adapter adapter that knows how to sent messages using different providers
type Adapter interface {
	SendSMS(ctx context.Context, to, message, from string) error
}
