package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	Logger interface {
		Info(msg string, args ...any)
		Error(msg string, args ...any)
	}

	// TelegramBotFetcher works with exact bot and exact token.
	TelegramBotFetcher interface {
		MakeRequest(endpoint string, params tgbotapi.Params) (*tgbotapi.APIResponse, error)
	}
)
