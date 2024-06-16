package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BusinessMessageConfig struct {
	tgbotapi.BaseChat
	tgbotapi.MessageConfig
	BusinessConnectionID string
}

func (config BusinessMessageConfig) params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)

	err := params.AddFirstValid("chat_id", config.BaseChat.ChatID, config.BaseChat.ChannelUsername)
	if err != nil {
		return nil, fmt.Errorf("addFirstValid: %w", err)
	}

	params.AddNonZero("reply_to_message_id", config.BaseChat.ReplyToMessageID)

	params.AddBool("disable_notification", config.BaseChat.DisableNotification)
	params.AddBool("allow_sending_without_reply", config.BaseChat.AllowSendingWithoutReply)

	err = params.AddInterface("reply_markup", config.BaseChat.ReplyMarkup)
	if err != nil {
		return nil, fmt.Errorf("add reply markup: %w", err)
	}

	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)

	params.AddNonEmpty("text", config.Text)
	params.AddBool("disable_web_page_preview", config.DisableWebPagePreview)
	params.AddNonEmpty("parse_mode", config.ParseMode)

	err = params.AddInterface("entities", config.Entities)
	if err != nil {
		return nil, fmt.Errorf("add entities: %w", err)
	}

	return params, nil
}

func (config BusinessMessageConfig) method() string {
	return "sendMessage"
}
