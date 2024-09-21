package handler

import (
	"fmt"
	"log/slog"

	"github.com/iooojik/tg-auto-response/internal/model"
)

type (
	Decision func(msg *model.BusinessMessageConfig) error

	Handler func(upd model.Update) error
)

func CheckIgnore(from model.IgnoreFrom) Handler {
	return func(upd model.Update) error {
		message := upd.BusinessMessage

		if from.Contains(message.From.ID) {
			return fmt.Errorf("%w: %v", ErrIgnore, message.From.ID)
		}

		return nil
	}
}

func DebugMessage(l Logger, debug bool) Handler {
	return func(upd model.Update) error {
		msg := upd.BusinessMessage

		if !debug || l == nil {
			return nil
		}

		slog.Info("msg", "chatID", msg.From.ID, "username", msg.From.UserName, "msg", msg.Text)

		return nil
	}
}

func HandleBusinessMessage(
	decision Decision,
	conditions ...model.Condition,
) Handler {
	return func(upd model.Update) error {
		for _, condition := range conditions {
			msg, err := CheckMessage(upd.BusinessMessage, condition)
			if err != nil {
				return fmt.Errorf("%w. condition reply from config: %v", err, condition.Reply)
			}

			if msg == nil {
				continue
			}

			err = decision(msg)
			if err != nil {
				return fmt.Errorf("decision: %w", err)
			}

			return nil
		}

		return nil
	}
}

func SendResponse(b TelegramBotFetcher) Decision {
	return func(msg *model.BusinessMessageConfig) error {
		params, err := msg.Params()
		if err != nil {
			return fmt.Errorf("params %w: %v", err, msg)
		}

		_, err = b.MakeRequest(msg.Method(), params)
		if err != nil {
			return fmt.Errorf("send message %w: %v", err, msg)
		}

		return nil
	}
}
