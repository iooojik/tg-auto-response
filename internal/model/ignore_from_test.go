package model_test

import (
	"testing"

	"github.com/iooojik/tg-auto-response/internal/model"
)

func TestIgnoreFrom_Contains(t *testing.T) {
	t.Parallel()

	type args struct {
		userID int64
	}

	tests := []struct {
		name string
		i    model.IgnoreFrom
		args args
		want bool
	}{
		{
			name: "test ignore",
			i:    model.IgnoreFrom{1234},
			args: args{1234},
			want: true,
		},
		{
			name: "test does not ignore",
			i:    model.IgnoreFrom{1234},
			args: args{9999},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.i.Contains(tt.args.userID); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
