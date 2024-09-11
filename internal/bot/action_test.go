package bot

import "testing"

func Test_sanitizeMessageText(t *testing.T) {
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
			if got := sanitizeMessageText(tt.args.s); got != tt.want {
				t.Errorf("sanitizeMessageText() = %v, want %v", got, tt.want)
			}
		})
	}
}
