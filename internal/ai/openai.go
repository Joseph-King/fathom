package ai

import "context"

type OpenAIProvider struct {
	BaseURL string
	Model   string
}

func NewOpenAIProvider(baseURL, model string) *OpenAIProvider {
	return &OpenAIProvider{BaseURL: baseURL, Model: model}
}

func (o *OpenAIProvider) Summary(ctx context.Context, logs string) (string, error) {
	// Your V1 logic using Ollama's local HTTP API goes here
	return "OpenAI analysis...", nil
}

func (o *OpenAIProvider) StreamSummary(ctx context.Context, logs string, callback func(string)) error {
	// Your V1 logic streaming local tokens goes here
	return nil
}
