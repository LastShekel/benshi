package worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getAvailablePort(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				"8080",
			},
			want: "8080",
		},
		{
			name: "change port",
			args: args{
				"-1",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getAvailablePort(tt.args.port), "getAvailablePort(%v)", tt.args.port)
		})
	}
}
