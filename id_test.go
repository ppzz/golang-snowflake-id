package id

import (
	"testing"
	"time"
)

func TestGetTimestamp_GetServerID_GetCounter(t *testing.T) {
	type args struct {
		ts  int64
		sid int64
		c   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1:普通参数",
			args: args{
				ts:  time.Now().UnixNano() / 1e6,
				sid: 2011,
				c:   1001,
			},
		},
		{
			name: "case2:全0参数:0,0,0",
			args: args{
				ts:  0,
				sid: 0,
				c:   0,
			},
		},
		{
			name: "case3:全MAX参数",
			args: args{
				ts:  0x3FFFFFFFFFF,
				sid: 0x7FF,
				c:   0x3FF,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := assemble(tt.args.ts, tt.args.sid, tt.args.c)
			ts := id.GetTimestamp()
			sid := id.GetServerID()
			c := id.GetCounter()
			if ts != tt.args.ts {
				t.Errorf("GetTimestamp() = %v, want %v", ts, tt.args.ts)
			}

			if sid != tt.args.sid {
				t.Errorf("GetServerID() = %v, want %v", sid, tt.args.sid)
			}
			if c != tt.args.c {
				t.Errorf("GetCounter() = %v, want %v", c, tt.args.c)
			}
		})
	}

}
