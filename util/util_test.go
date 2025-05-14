package util_test

import (
	"testing"

	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/util"
)

func TestGetDataKeyFromEventKey(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	type args struct {
		eventKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				eventKey: "time2go:trigger:clientName:eventName:eventId-14",
			},
			want: "time2go:data:clientName:eventName:eventId-14",
		},
		{
			name: "false",
			args: args{
				eventKey: "time2go:no:clientName:eventName:eventId-14",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.GetDataKeyFromEventKey(tt.args.eventKey); got != tt.want {
				t.Errorf("GetDataKeyFromEventKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
