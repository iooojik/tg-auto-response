package bot

import (
	"regexp"
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	spaceRegexp = regexp.MustCompile(`\s{2,}`)
)

func actionCheckHello(
	message *BusinessMessage,
	cfg *AutoHello,
) (*BusinessMessageConfig, error) {
	if cfg == nil {
		return nil, nil
	}

	content := ExtractTextAndSpaces(strings.ToLower(strings.TrimSpace(message.Text)))

	for _, hello := range cfg.IncomeMessages {
		if !strings.EqualFold(strings.ToLower(hello), content) {
			continue
		}

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

	return nil, nil
}

func ExtractTextAndSpaces(s string) string {
	var builder strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			builder.WriteRune(char)
		}
	}

	s = strings.Trim(builder.String(), "\n\t\r ")

	s = spaceRegexp.ReplaceAllString(s, " ")

	return s
}
