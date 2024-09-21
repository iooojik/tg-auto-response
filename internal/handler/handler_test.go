//nolint:exhaustruct
package handler_test

import (
	"errors"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iooojik/tg-auto-response/internal/handler"
	"github.com/iooojik/tg-auto-response/internal/model"
)

func TestCheckIgnore(t *testing.T) {
	t.Parallel()

	type args struct {
		from model.IgnoreFrom
		upd  model.Update
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "test does not ignore",
			args: args{
				from: model.IgnoreFrom{
					123456,
				},
				upd: model.Update{
					BusinessMessage: &model.BusinessMessage{
						Message: &tgbotapi.Message{
							From: &tgbotapi.User{ID: 9999},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "test ignore",
			args: args{
				from: model.IgnoreFrom{
					123456,
				},
				upd: model.Update{
					BusinessMessage: &model.BusinessMessage{
						Message: &tgbotapi.Message{
							From: &tgbotapi.User{ID: 123456},
						},
					},
				},
			},
			wantErr: handler.ErrIgnore,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotFunc := handler.CheckIgnore(tt.args.from)

			err := gotFunc(tt.args.upd)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CheckIgnore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
