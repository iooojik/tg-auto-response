package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type HandleCfg struct {
	IncomeMessages []string `yaml:"income_messages"`
	Reply          string   `yaml:"reply"`
}

type Config struct {
	Token              string      `yaml:"token"`
	Debug              bool        `yaml:"debug"`
	Handle             []HandleCfg `yaml:"handle"`
	IgnoreMessagesFrom []int64     `yaml:"ignore_messages_from"`
}

type BusinessMessage struct {
	BusinessConnectionID string `json:"business_connection_id"`
	*tgbotapi.Message
}

type Update struct {
	tgbotapi.Update
	BusinessMessage *BusinessMessage `json:"business_message"`
}
