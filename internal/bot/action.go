package bot

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	spaceRegexp = regexp.MustCompile(`\s{2,}`)
)

func CheckMessage(
	message *BusinessMessage,
	cfg HandleCfg,
) (*BusinessMessageConfig, error) {
	if len(cfg.IncomeMessages) == 0 || cfg.Reply == "" {
		return nil, errors.New("configuration is nil")
	}

	content := sanitizeMessageText(message.Text)

	for _, hello := range cfg.IncomeMessages {
		if !strings.EqualFold(hello, content) {
			continue
		}

		msg := &BusinessMessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           message.Chat.ID,
				ReplyToMessageID: 0,
			},
			MessageConfig: tgbotapi.MessageConfig{
				Text:                  cfg.Reply,
				ParseMode:             "MarkdownV2",
				DisableWebPagePreview: true,
			},
			BusinessConnectionID: message.BusinessConnectionID,
		}

		return msg, nil
	}

	return nil, nil
}

func sanitizeMessageText(s string) string {
	var result []rune

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			result = append(result, char)
		}
	}

	s = strings.TrimSpace(string(result))
	s = spaceRegexp.ReplaceAllString(s, " ")

	return s
}
