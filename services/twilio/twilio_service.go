package twilio

import (
	"log"

	"github.com/twilio/twilio-go"
	lookups "github.com/twilio/twilio-go/rest/lookups/v2"
)

var Client *twilio.RestClient

func InitClient(accountSid, authToken string) {
	Client = twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: accountSid,
		Password:   authToken,
	})
}

func GetClient() *twilio.RestClient {
	return Client
}

func TestClient(number string) error {
	log.Printf("running twilio client test...")

	params := &lookups.FetchPhoneNumberParams{}
	resp, err := Client.LookupsV2.FetchPhoneNumber(number, params)
	if err != nil {
		log.Fatalf("error fetching phone number: %s", err)
		return err
	}

	log.Printf("Phone number: %s", *resp.PhoneNumber)
	return nil
}
