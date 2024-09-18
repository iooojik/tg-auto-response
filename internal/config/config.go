package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/iooojik/tg-auto-response/internal/bot"
)

type Config struct {
	BotConfig *bot.Config `yaml:"bot"`
}

func ReadCfg(p string) (*Config, error) {
	f, err := os.OpenFile(p, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", p, err)
	}

	cfg := new(Config)

	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config %s: %w", p, err)
	}

	return cfg, nil
}
