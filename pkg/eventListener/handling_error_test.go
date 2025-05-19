package eventlistener

import (
	"testing"
	"time"
)

func Test_exponentialTime(t *testing.T) {
	type args struct {
		baseTime     time.Duration
		retryCounter int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "success",
			args: args{
				baseTime:     time.Second,
				retryCounter: 3,
			},
			want: 8 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exponentialTime(tt.args.baseTime, tt.args.retryCounter); got != tt.want {
				t.Errorf("exponentialTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
