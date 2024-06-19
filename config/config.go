package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	OpenAI struct {
		ApiKey string `envconfig:"OPENAI_API_KEY"`
	}
	Twilio struct {
		AccountSid string `envconfig:"TWILIO_ACCOUNT_SID"`
		AuthToken  string `envconfig:"TWILIO_AUTH_TOKEN"`
	}
}

func NewConfig() (*Config, error) {
	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
