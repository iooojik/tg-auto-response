package bot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iooojik/tg-auto-response/internal/handler"
	"github.com/iooojik/tg-auto-response/internal/model"
)

const (
	ReceiveMessagesTimeout = 5

	UpdateBuffSize = 1
)

type (
	UpdatesHandler func(updates *chan model.Update) error
)

type Bot struct {
	// chat id - command
	chatContext     map[int64]string
	shutdownChannel chan any
	botAPI          BotAPI
	updatesHandler  UpdatesHandler
}

// New Creates a new Bot instance.
func New(
	cfg *Config,
	logger Logger,
) *Bot {
	bot, err := authorizeBot(cfg.Token, cfg.Debug)
	if err != nil {
		panic(fmt.Errorf("authorize bot: %w", err))
	}

	b := &Bot{
		chatContext: map[int64]string{},
		botAPI:      bot,
		updatesHandler: handleUpdates(
			logger,
			handler.DebugMessage(logger, cfg.Debug),
			handler.CheckIgnore(cfg.IgnoreMessagesFrom),
			handler.HandleBusinessMessage(
				handler.SendResponse(bot),
				cfg.Conditions...,
			),
		),
	}

	return b
}

func (b *Bot) Run(
	ctx context.Context,
) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = ReceiveMessagesTimeout

	updates := b.GetUpdates(ctx, b.botAPI, u)

	err := b.updatesHandler(&updates)
	if err != nil {
		return fmt.Errorf("handle updates: %w", err)
	}

	return nil
}

// GetUpdates starts and returns a channel for getting updates.
func (b *Bot) GetUpdates(
	ctx context.Context,
	tgAPI TelegramFetcher,
	config tgbotapi.UpdateConfig,
) chan model.Update {
	ch := make(chan model.Update, UpdateBuffSize)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
			case <-b.shutdownChannel:
				close(ch)
				return

			default:
			}

			updates, err := FetchUpdates(tgAPI, config)
			if err != nil {
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
	}(ctx)

	return ch
}

func FetchUpdates(tgAPI TelegramFetcher, config tgbotapi.UpdateConfig) ([]model.Update, error) {
	resp, err := tgAPI.Request(config)
	if err != nil {
		return make([]model.Update, 0), fmt.Errorf("fetch updates: %w", err)
	}

	var updates []model.Update

	err = json.Unmarshal(resp.Result, &updates)
	if err != nil {
		return make([]model.Update, 0), fmt.Errorf("unmarshal updates: %w", err)
	}

	return updates, nil
}

func authorizeBot(
	token string,
	debug bool,
) (BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("new bot: %w", err)
	}

	bot.Debug = debug

	return bot, nil
}

func handleUpdates(
	l Logger,
	handlers ...handler.Handler,
) UpdatesHandler {
	return func(updates *chan model.Update) error {
		for update := range *updates {
			if update.BusinessMessage == nil {
				continue
			}

			for _, h := range handlers {
				err := h(update)
				if err != nil {
					l.Error("handle", "msg", err.Error())
				}

				if errors.Is(err, handler.ErrIgnore) {
					break
				}
			}
		}

		return nil
	}
}
