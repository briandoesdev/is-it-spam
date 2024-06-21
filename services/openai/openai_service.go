package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/briandoesdev/caller-lookup/config"
)

var (
	Service                 = &OpenAIService{}
	openaiBaseUrl           = "https://api.openai.com/v1/chat/completions"
	errUninitializedService = errors.New("openai service not initialized")
)

var (
	httpTransport *http.Transport
	httpClient    *http.Client
)

type OpenAIService struct {
	Model     string
	ApiKey    string
	OrgID     string
	ProjectID string
}

type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

func InitService(c config.OpenAI) {
	log.Printf("Initializing OpenAI service.")

	Service.ApiKey = c.ApiKey
	Service.Model = c.Model
	Service.OrgID = c.OrganizationID
	Service.ProjectID = c.ProjectID

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
	message := OpenAIRequest{
		Model: Service.Model,
		Messages: []OpenAIMessage{
			{Role: "system", Content: "Summarize the JSON blob given by the user. Format your response in a paragraph to include the following details: number user and owner, current address, related individuals, alternate names, and alternate numbers. Starts your response like so: \"{phone number} belongs to \". End each response with whether the number is likely spam or not. A number is likely spam if it has no associated owner name or is VOIP. "},
			{Role: "user", Content: prompt},
		},
		Temperature: 0.7,
	}
	messagejson, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	bodyReader := bytes.NewReader(messagejson)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Service.ApiKey)
	req.Header.Add("OpenAI-Organization", Service.OrgID)
	req.Header.Add("OpenAI-Project", Service.ProjectID)

	r, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("httpclient error: %w", err)
	}
	defer io.Copy(io.Discard, r.Body)
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return "", fmt.Errorf("generate completions failed: %s", r.Status)
	}

	var result OpenAIResponse
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return "", err
	}

	return strings.Trim(result.Choices[0].Message.Content, "\""), nil
}

func checkInit() error {
	if Service.ApiKey == "" {
		return errUninitializedService
	}
	return nil
}
