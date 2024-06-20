package twilio

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	Service                 = &TwilioService{}
	twilioBaseUrl           = "https://lookups.twilio.com/v1/PhoneNumbers/"
	twilioAddOns            = "?AddOns=trestle_reverse_phone"
	errUninitializedService = errors.New("openai service not initialized")
	errLookupFailed         = errors.New("lookup failed")
)

var (
	httpTransport *http.Transport
	httpClient    *http.Client
)

type TwilioService struct {
	accountSid  string
	authToken   string
	EncodedAuth string
}

func InitService(accountSid, authToken string) {
	log.Printf("Initializing Twilio service.")

	Service.accountSid = accountSid
	Service.authToken = authToken
	Service.EncodedAuth = base64.StdEncoding.EncodeToString([]byte(accountSid + ":" + authToken))

	httpTransport = &http.Transport{
		MaxIdleConns:      10,
		IdleConnTimeout:   15 * time.Second,
		DisableKeepAlives: false,
	}

	httpClient = &http.Client{
		Transport: httpTransport,
		Timeout:   30 * time.Second,
	}
}

func GetService() (*TwilioService, error) {
	if err := checkInit(); err != nil {
		return nil, err
	}

	return Service, nil
}
func GetClient() *TwilioService {
	return Service
}

func Lookup(number string) (string, error) {
	if err := checkInit(); err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s%s%s", twilioBaseUrl, number, twilioAddOns)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+Service.EncodedAuth)

	r, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer io.Copy(io.Discard, r.Body)
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return "", errLookupFailed
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errLookupFailed
	}

	return string(b), nil
}

func checkInit() error {
	if Service.accountSid == "" || Service.authToken == "" {
		return errUninitializedService
	}
	return nil
}
