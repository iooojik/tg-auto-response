package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os/signal"
	"strings"
	"syscall"

	"github.com/iooojik/tg-auto-response/internal/bot"
	"github.com/iooojik/tg-auto-response/internal/config"
)

func main() {
	cfg := flag.String("cfg", "configs/config.yaml", "app config")
	flag.Parse()

	if err := runBot(*cfg); err != nil {
		panic(err)
	}
}

func runBot(cfg string) error {
	logger := slog.Default()

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	defer stop()

	logger.Info("config", "loading from", cfg)

	conf, err := config.ReadCfg(cfg)
	if err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}

	b := bot.New(conf.BotConfig, logger)

	logger.Info("starting", "bot_login", strings.Split(conf.BotConfig.Token, ":")[0])

	if err = b.Run(ctx); err != nil {
		return fmt.Errorf("failed to run bot: %w", err)
	}

	logger.Info("stopped", "bot_login", strings.Split(conf.BotConfig.Token, ":")[0])

	return nil
}
