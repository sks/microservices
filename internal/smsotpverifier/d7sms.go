package smsotpverifier

import (
	context "context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ErrMissingRequiredCreds some of the required credentials were missing at runtime
var ErrMissingRequiredCreds = errors.New("username/password are required parameter to use D7 Network API")

/*
D7SMS a struct that knows how to talk to D7SMS
REST API docs can be found at https://d7networks.com/docs/apis/rest/index.html
**/
type D7SMS struct {
	authHeader string
}

const sendAPI = "/secure/send"
const d7networkBaseURI = "https://rest-api.d7networks.com"

// NewD7SMSProvider create the sms provider
func NewD7SMSProvider(username, password string) (Adapter, error) {
	if username == "" || password == "" {
		return nil, ErrMissingRequiredCreds
	}
	authHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return &D7SMS{
		authHeader,
	}, nil
}

// SendSMS send sms to the given number
func (d *D7SMS) SendSMS(ctx context.Context, to, message, from string) error {

	payload := fmt.Sprintf(`{"to": %s,"from": %q,"coding":"8","hex-content": %q}`, to, from, hex.EncodeToString([]byte(message)))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", d7networkBaseURI, sendAPI), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", d.authHeader))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status %d, got %d", http.StatusOK, res.StatusCode)
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}
