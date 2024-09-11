package bot

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	cfg *Config
	// chat id - command
	chatContext     map[int64]string
	Buffer          int
	shutdownChannel chan any
	botAPI          *tgbotapi.BotAPI
}

func New(
	cfg *Config,
) *Bot {
	b := &Bot{
		cfg:         cfg,
		chatContext: map[int64]string{},
		Buffer:      1,
	}

	return b
}

func (b *Bot) Run() error {
	bot, err := tgbotapi.NewBotAPI(b.cfg.Token)
	if err != nil {
		return fmt.Errorf("new bot: %w", err)
	}

	b.botAPI = bot
	b.botAPI.Debug = b.cfg.Debug

	slog.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates := b.GetUpdatesChan(u)

	for update := range updates {
		var (
			msg *BusinessMessageConfig
			err error
		)

		if update.BusinessMessage != nil {
			if b.isIgnore(update.BusinessMessage.From.ID) {
				continue
			}

			msg, err = b.handleBusinessMessage(update.BusinessMessage)
		} else {
			continue
		}

		if err != nil {
			slog.Error("handle", "err", err.Error())
			continue
		}

		if msg == nil {
			continue
		}

		params, err := msg.params()
		if err != nil {
			slog.Error("params", "err", err.Error())
			continue
		}

		_, err = bot.MakeRequest(msg.method(), params)
		if err != nil {
			slog.Error("send", "err", err.Error())
		}
	}

	return nil
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (b *Bot) GetUpdatesChan(config tgbotapi.UpdateConfig) chan Update {
	ch := make(chan Update, b.Buffer)

	go func() {
		for {
			select {
			case <-b.shutdownChannel:
				close(ch)
				return

			default:
			}

			updates, err := b.GetUpdates(config)
			if err != nil {
				slog.Error("updates", "err", err.Error())
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (b *Bot) isIgnore(userID int64) bool {
	for _, v := range b.cfg.IgnoreMessagesFrom {
		if v == userID {
			return true
		}
	}

	return false
}

func (b *Bot) GetUpdates(config tgbotapi.UpdateConfig) ([]Update, error) {
	resp, err := b.botAPI.Request(config)
	if err != nil {
		return []Update{}, err
	}

	var updates []Update
	err = json.Unmarshal(resp.Result, &updates)

	return updates, err
}

func (b *Bot) handleBusinessMessage(
	message *BusinessMessage,
) (*BusinessMessageConfig, error) {
	if b.cfg.Debug {
		slog.Info("msg", "chat_ID", message.From.ID, "username", message.From.UserName, "msg", message.Text)
	}

	for _, handleCfg := range b.cfg.Handle {
		msg, err := CheckMessage(message, handleCfg)
		if err != nil {
			slog.Error("action", "err", err.Error())
		}

		if msg != nil {
			return msg, nil
		}
	}

	return nil, nil
}
