package bot

import (
	"github.com/iooojik/tg-auto-response/internal/model"
)

type Config struct {
	Token              string            `yaml:"token"`
	Debug              bool              `yaml:"debug"`
	Conditions         []model.Condition `yaml:"conditions"`
	IgnoreMessagesFrom model.IgnoreFrom  `yaml:"ignore_messages_from"`
}
