//nolint:exhaustruct
package handler

import (
	"reflect"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iooojik/tg-auto-response/internal/model"
)

func Test_sanitizeMessageText(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test#1",
			args: args{
				s: "Привет 👋",
			},
			want: "Привет",
		},
		{
			name: "test#2",
			args: args{
				s: "Привет 👋 Как дела?",
			},
			want: "Привет Как дела",
		},
		{
			name: "test#3",
			args: args{
				s: "Привет 👋 Как дела?  \n   Как ты?",
			},
			want: "Привет Как дела Как ты",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := sanitizeMessage(tt.args.s)

			if got != tt.want {
				t.Errorf("sanitizeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckMessage(t *testing.T) {
	t.Parallel()

	type args struct {
		message   *model.BusinessMessage
		condition model.Condition
	}

	tests := []struct {
		name    string
		args    args
		want    *model.BusinessMessageConfig
		wantErr bool
	}{
		{
			name: "test valid",
			args: args{
				message: &model.BusinessMessage{
					BusinessConnectionID: "conn_id",
					Message: &tgbotapi.Message{
						Text: "how are you?",
						Chat: &tgbotapi.Chat{ID: 222},
					},
				},
				condition: model.Condition{
					Reply: "Nice! And you?",
					IncomeMessages: []string{
						"how are you",
					},
				},
			},
			want: &model.BusinessMessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: 222,
				},
				MessageConfig: tgbotapi.MessageConfig{
					Text:                  "Nice! And you?",
					ParseMode:             tgbotapi.ModeMarkdownV2,
					DisableWebPagePreview: true,
				},
				BusinessConnectionID: "conn_id",
			},
			wantErr: false,
		},
		{
			name:    "test without condition",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test without message",
			args: args{
				condition: model.Condition{
					Reply:          "Nice! And you?",
					IncomeMessages: []string{"how are you"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CheckMessage(tt.args.message, tt.args.condition)

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
