package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	OpenAI OpenAI
	Twilio Twilio
}

type OpenAI struct {
	ApiKey         string `envconfig:"OPENAI_API_KEY"`
	Model          string `envconfig:"OPENAI_MODEL"`
	OrganizationID string `envconfig:"OPENAI_ORGANIZATION_ID"`
	ProjectID      string `envconfig:"OPENAI_PROJECT_ID"`
}

type Twilio struct {
	AccountSid string `envconfig:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `envconfig:"TWILIO_AUTH_TOKEN"`
}

func NewConfig() (*Config, error) {
	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
