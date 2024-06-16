package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func actionCheckHello(
	message *BusinessMessage,
	cfg *AutoHello,
) (*BusinessMessageConfig, error) {
	if cfg == nil {
		return nil, nil
	}

	content := strings.ToLower(strings.TrimSpace(message.Text))

	for _, hello := range cfg.IncomeMessages {
		if strings.EqualFold(strings.ToLower(hello), content) {
			msg := BusinessMessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID:           message.Chat.ID,
					ReplyToMessageID: 0,
				},
				MessageConfig: tgbotapi.MessageConfig{
					Text:                  cfg.Reply,
					DisableWebPagePreview: true,
				},
				BusinessConnectionID: message.BusinessConnectionID,
			}

			return &msg, nil
		}
	}

	return nil, nil
}
