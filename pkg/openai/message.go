package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (c *Client) GenerateResponse(ctx context.Context, in string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: in,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("chat completion: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
