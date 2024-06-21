package twilio

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/briandoesdev/caller-lookup/config"
)

var (
	Service                 = &TwilioService{}
	twilioBaseUrl           = "https://lookups.twilio.com/v1/PhoneNumbers/"
	twilioAddOns            = "?AddOns=trestle_reverse_phone"
	errUninitializedService = errors.New("openai service not initialized")
	errLookupFailed         = errors.New("error looking up number, verify the number and try again")
)

var (
	httpTransport *http.Transport
	httpClient    *http.Client
)

type TwilioService struct {
	accountSid string
	authToken  string
}

func InitService(c config.Twilio) {
	log.Printf("Initializing Twilio service.")

	Service.accountSid = c.AccountSid
	Service.authToken = c.AuthToken

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

	req.SetBasicAuth(Service.accountSid, Service.authToken)

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
