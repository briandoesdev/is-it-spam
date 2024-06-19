package twilio

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	Client        = &TwilioClient{}
	twilioBaseUrl = "https://lookups.twilio.com/v1/PhoneNumbers/"
	twilioAddOns  = "?AddOns=trestle_reverse_phone"
)

type TwilioClient struct {
	accountSid  string
	authToken   string
	EncodedAuth string
}

func InitService(accountSid, authToken string) {
	Client.accountSid = accountSid
	Client.authToken = authToken
	Client.EncodedAuth = base64.StdEncoding.EncodeToString([]byte(accountSid + ":" + authToken))
}

func GetClient() *TwilioClient {
	return Client
}

func Lookup(number string) (string, error) {
	if err := checkInit(); err != nil {
		return "", fmt.Errorf("twilio client not initialized")
	}

	url := fmt.Sprintf("%s%s%s", twilioBaseUrl, number, twilioAddOns)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+Client.EncodedAuth)

	h := &http.Client{Timeout: 10 * time.Second}
	r, err := h.Do(req)

	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return "", err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func checkInit() error {
	if Client.accountSid == "" || Client.authToken == "" {
		return fmt.Errorf("twilio client not initialized")
	}
	return nil
}
