package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/iooojik/tg-auto-hello/internal/bot"
	"github.com/iooojik/tg-auto-hello/internal/config"
)

func main() {
	cfg := flag.String("cfg", "configs/config.yaml", "app config")
	flag.Parse()

	if err := runBot(*cfg); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}
}

func runBot(cfg string) error {
	log.Printf("Loading configuration from %s", cfg)
	conf, err := config.ReadCfg(cfg)
	if err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}

	b := bot.New(conf.BotConfig)

	log.Println("Starting the bot...")
	if err := b.Run(); err != nil {
		return fmt.Errorf("failed to run bot: %w", err)
	}

	log.Println("Bot stopped successfully.")
	return nil
}
