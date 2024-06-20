package openai

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	Service                 = &OpenAIService{}
	openaiBaseUrl           = "https://api.openai.com/v1/chat/completions"
	errUninitializedService = errors.New("openai service not initialized")
	errGenerateFailed       = errors.New("generate completions failed")
)

var (
	httpTransport *http.Transport
	httpClient    *http.Client
)

type OpenAIService struct {
	Model  string
	ApiKey string
}

func InitService(apiKey, model string) {
	log.Printf("Initializing OpenAI service.")

	Service.ApiKey = apiKey
	Service.Model = model

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

func GetService() (*OpenAIService, error) {
	if err := checkInit(); err != nil {
		return nil, err
	}

	return Service, nil
}

func GenerateCompletions(prompt string) (string, error) {
	if err := checkInit(); err != nil {
		return "", err
	}

	url := openaiBaseUrl
	body := []byte(`{"model": "` + Service.Model + `", "messages": [{"role": "user", "content": "` + prompt + `"}]}`)
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Service.ApiKey)

	r, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer io.Copy(io.Discard, r.Body)
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return "", fmt.Errorf("generate completions failed: %s", r.Status)
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func checkInit() error {
	if Service.ApiKey == "" {
		return errUninitializedService
	}
	return nil
}
