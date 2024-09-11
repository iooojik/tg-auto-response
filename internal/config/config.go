package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/iooojik/tg-auto-hello/internal/bot"
)

type Config struct {
	BotConfig *bot.Config `yaml:"bot"`
}

func ReadCfg(p string) (*Config, error) {
	log.Printf("Reading configuration file: %s\n", p)
	f, err := os.OpenFile(p, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", p, err)
	}

	cfg := new(Config)

	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config %s: %w", p, err)
	}

	log.Println("Configuration file loaded successfully.")
	return cfg, nil
}
