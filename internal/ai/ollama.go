package ai

import "context"

type OllamaProvider struct {
	BaseURL string
	Model   string
}

func NewOllamaProvider(baseURL, model string) *OllamaProvider {
	return &OllamaProvider{BaseURL: baseURL, Model: model}
}

func (o *OllamaProvider) Summary(ctx context.Context, logs string) (string, error) {
	// Your V1 logic using Ollama's local HTTP API goes here
	return "Ollama analysis...", nil
}

func (o *OllamaProvider) StreamSummary(ctx context.Context, logs string, callback func(string)) error {
	// Your V1 logic streaming local tokens goes here
	return nil
}
