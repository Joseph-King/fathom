// Copyright 2026 Joseph King
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
