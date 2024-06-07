package model

import (
	"reflect"
	"testing"
)

func TestParseColonField(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Via: SIP/2.0/UDP 100.85.93.61:48185;rport=48185;received=100.85.93.61;branch=zzzzzzzzzzzzz\r\nCall-ID: 1234567890",
			args: args{
				str: "Via: SIP/2.0/UDP 100.85.93.61:48185;rport=48185;received=100.85.93.61;branch=zzzzzzzzzzzzz\r\nCall-ID: 1234567890",
			},
			want: map[string]string{
				"Via":     "SIP/2.0/UDP 100.85.93.61:48185;rport=48185;received=100.85.93.61;branch=zzzzzzzzzzzzz",
				"Call-ID": "1234567890",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseColonField(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseColonField() = %v, want %v", got, tt.want)
			}
		})
	}
}
