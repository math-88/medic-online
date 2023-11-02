package openai

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/math-88/medic-online/internal/config"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	config *config.Config
	client *openai.Client
}

func New(cfg *config.Config) *OpenAI {

	client := openai.NewClient(cfg.Openai.Token)

	return &OpenAI{
		config: cfg,
		client: client,
	}
}

func (oi *OpenAI) Ask(
	ctx context.Context,
	q string,
) (string, error) {

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: q,
			},
		},
	}

	resp, err := oi.client.CreateChatCompletion(
		ctx,
		req,
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func (oi *OpenAI) AskStream(
	ctx context.Context,
	q string,
	ch chan<- string,
) error {

	defer close(ch)

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: q,
			},
		},
		Stream: true,
	}
	stream, err := oi.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return fmt.Errorf("ChatCompletionStream error: %w", err)
	}
	defer stream.Close()

	if err != nil {
		return fmt.Errorf("ChatCompletion error: %w", err)
	}

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("stream error: %w", err)
		}

		ch <- response.Choices[0].Delta.Content
	}
}
