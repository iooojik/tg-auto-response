package main

import (
	"flag"
	"fmt"

	"auto-hello/internal/bot"
	"auto-hello/internal/config"
)

func main() {
	cfg := flag.String("cfg", "configs/config.yaml", "app config")
	flag.Parse()

	err := runBot(*cfg)
	if err != nil {
		panic(err)
	}
}

func runBot(cfg string) error {
	conf, err := config.ReadCfg(cfg)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	b := bot.NewBot(conf.BotConfig)

	err = b.Run()
	if err != nil {
		return fmt.Errorf("run bot: %w", err)
	}

	return nil
}
