package llms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OllamaClient is a client for the Ollama API
type OllamaClient struct {
	httpClient *http.Client
	baseURL    string
	model      string
}

// OllamaRequest represents the request body for Ollama API
type OllamaRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// OllamaResponse represents the response from Ollama API
type OllamaResponse struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	Context   []int  `json:"context,omitempty"`
	TotalTime int64  `json:"total_time,omitempty"`
}

// ClientOption is a function that configures an OllamaClient
type ClientOption func(*OllamaClient)

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *OllamaClient) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient allows setting a custom http.Client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *OllamaClient) {
		c.httpClient = client
	}
}

// NewOllamaClient creates a new OllamaClient instance with optional configuration
func NewOllamaClient(baseURL, model string, opts ...ClientOption) *OllamaClient {
	client := &OllamaClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		model:   model,
	}

	// Apply any provided options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Generate sends a prompt to Ollama and returns the response
func (o *OllamaClient) Generate(ctx context.Context, prompt string) (string, error) {
	req := OllamaRequest{
		Model:  o.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", o.baseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := o.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse ollama response: %w", err)
	}

	return result.Response, nil
}

// GenerateWithOptions sends a prompt to Ollama with custom options and returns the response
func (o *OllamaClient) GenerateWithOptions(ctx context.Context, prompt string, options map[string]interface{}) (string, error) {
	req := OllamaRequest{
		Model:   o.model,
		Prompt:  prompt,
		Stream:  false,
		Options: options,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", o.baseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := o.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse ollama response: %w", err)
	}

	return result.Response, nil
}

// SetModel changes the model used by the client
func (o *OllamaClient) SetModel(model string) {
	o.model = model
}

// SetTimeout sets the request timeout
func (o *OllamaClient) SetTimeout(timeout time.Duration) {
	o.httpClient.Timeout = timeout
}

// GetModel returns the current model name
func (o *OllamaClient) GetModel() string {
	return o.model
}

// GetBaseURL returns the current base URL
func (o *OllamaClient) GetBaseURL() string {
	return o.baseURL
}

