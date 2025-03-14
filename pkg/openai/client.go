package openai

import (
	"github.com/sashabaranov/go-openai"
)

type Config struct {
	Token string `yaml:"token"`
}

type Client struct {
	client *openai.Client
}

func New(cfg Config) *Client {
	cl := &Client{
		client: openai.NewClient(cfg.Token),
	}

	return cl
}
