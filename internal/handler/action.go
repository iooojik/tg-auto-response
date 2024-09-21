package handler

import (
	"regexp"
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iooojik/tg-auto-response/internal/model"
)

var (
	spaceRegexp = regexp.MustCompile(`\s{2,}`)
)

func CheckMessage(
	message *model.BusinessMessage,
	condition model.Condition,
) (*model.BusinessMessageConfig, error) {
	if len(condition.IncomeMessages) == 0 || condition.Reply == "" {
		return nil, ErrNoCondition
	}

	if message == nil {
		return nil, ErrNoMessage
	}

	content := sanitizeMessage(message.Text)

	for _, incomeMessage := range condition.IncomeMessages {
		if !strings.EqualFold(incomeMessage, content) {
			continue
		}

		//nolint:exhaustruct
		msg := &model.BusinessMessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           message.Chat.ID,
				ReplyToMessageID: 0,
			},
			MessageConfig: tgbotapi.MessageConfig{
				Text:                  condition.Reply,
				ParseMode:             tgbotapi.ModeMarkdownV2,
				DisableWebPagePreview: true,
			},
			BusinessConnectionID: message.BusinessConnectionID,
		}

		return msg, nil
	}

	//nolint:nilnil
	return nil, nil
}

func sanitizeMessage(s string) string {
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
