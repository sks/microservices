package smsotpverifier

import (
	context "context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/sks/microservices/internal/httputil"
	"go.uber.org/zap"
)

// ErrMissingRequiredCreds some of the required credentials were missing at runtime
var ErrMissingRequiredCreds = errors.New("username/password are required parameter to use D7 Network API")

/*
D7SMS a struct that knows how to talk to D7SMS
REST API docs can be found at https://d7networks.com/docs/apis/rest/index.html
**/
type D7SMS struct {
	authHeader string
	logger     *zap.Logger
}

const sendAPI = "secure/send"
const d7networkBaseURI = "https://rest-api.d7networks.com"

// NewD7SMSProvider create the sms provider
func NewD7SMSProvider(username, password string, logger *zap.Logger) (Adapter, error) {
	if username == "" || password == "" {
		return nil, ErrMissingRequiredCreds
	}
	authHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return &D7SMS{
		authHeader, logger,
	}, nil
}

// SendSMS send sms to the given number
func (d *D7SMS) SendSMS(ctx context.Context, to, message, from string) error {
	logger := d.logger.With(zap.String("to", to), zap.String("from", from))
	payload := fmt.Sprintf(`{"to": %s,"content": %q, "from":"SMSINFO"}`, to, message)
	logger.Info("Sending the sms", zap.Int("length", len(payload)))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", d7networkBaseURI, sendAPI), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", d.authHeader))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")

	body, statusCode, err := httputil.GetResponse(req)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.New(string(body))
	}
	logger.Debug("Done with sending sms", zap.String("Response", string(body)))
	return nil
}
