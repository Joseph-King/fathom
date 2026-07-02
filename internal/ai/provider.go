package ai

import "context"

// Provider defines the capabilities any LLM backend must have to work with Fathom.
type Provider interface {
	// Summary takes a raw wall of logs and returns a human-readable analysis.
	Summary(ctx context.Context, logs string) (string, error)

	// StreamSummary does the same, but yields tokens in real-time to the terminal.
	StreamSummary(ctx context.Context, logs string, callback func(string)) error
}
