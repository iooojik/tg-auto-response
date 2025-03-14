package bot

import (
	"github.com/iooojik/tg-auto-response/internal/model"
	"github.com/iooojik/tg-auto-response/pkg/openai"
)

type Config struct {
	Token              string            `yaml:"token"`
	Debug              bool              `yaml:"debug"`
	OpenAI             openai.Config     `yaml:"openai"`
	Conditions         []model.Condition `yaml:"conditions"`
	IgnoreMessagesFrom model.IgnoreFrom  `yaml:"ignore_messages_from"`
}
