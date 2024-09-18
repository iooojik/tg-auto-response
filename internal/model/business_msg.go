package model

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BusinessMessage struct {
	BusinessConnectionID string `json:"business_connection_id"`
	*tgbotapi.Message
}

type Update struct {
	tgbotapi.Update
	BusinessMessage *BusinessMessage `json:"business_message"`
}

type BusinessMessageConfig struct {
	tgbotapi.BaseChat
	tgbotapi.MessageConfig
	BusinessConnectionID string
}

func (config BusinessMessageConfig) Params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)

	err := addFirstValidParam(params, "chat_id", config.BaseChat.ChatID, config.BaseChat.ChannelUsername)
	if err != nil {
		return nil, err
	}

	params.AddNonZero("reply_to_message_id", config.BaseChat.ReplyToMessageID)
	params.AddBool("disable_notification", config.BaseChat.DisableNotification)
	params.AddBool("allow_sending_without_reply", config.BaseChat.AllowSendingWithoutReply)

	err = addInterfaceParam(params, "reply_markup", config.BaseChat.ReplyMarkup)
	if err != nil {
		return nil, err
	}

	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("text", config.Text)
	params.AddBool("disable_web_page_preview", config.DisableWebPagePreview)
	params.AddNonEmpty("parse_mode", config.ParseMode)

	err = addInterfaceParam(params, "entities", config.Entities)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (config BusinessMessageConfig) Method() string {
	return "sendMessage"
}

// Helper functions to simplify repetitive parameter addition and error handling.
func addFirstValidParam(params tgbotapi.Params, key string, values ...interface{}) error {
	err := params.AddFirstValid(key, values...)
	if err != nil {
		return fmt.Errorf("failed to add parameter %s: %w", key, err)
	}

	return nil
}

func addInterfaceParam(params tgbotapi.Params, key string, value interface{}) error {
	err := params.AddInterface(key, value)
	if err != nil {
		return fmt.Errorf("failed to add interface parameter %s: %w", key, err)
	}

	return nil
}
