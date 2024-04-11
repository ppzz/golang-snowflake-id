package id

import (
	"testing"
	"time"
)

func TestGetTimestamp_GetServerID_GetCounter(t *testing.T) {
	type args struct {
		ts  uint64
		sid uint64
		c   uint64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1:普通参数",
			args: args{
				ts:  uint64(time.Now().UnixNano() / 1e6),
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
			if ts != int64(tt.args.ts) {
				t.Errorf("GetTimestamp() = %v, id %v", ts, tt.args.ts)
			}

			if sid != int64(tt.args.sid) {
				t.Errorf("GetServerID() = %v, id %v", sid, tt.args.sid)
			}
			if c != int64(tt.args.c) {
				t.Errorf("GetCounter() = %v, id %v", c, tt.args.c)
			}
		})
	}

}

func TestID_IdStr(t *testing.T) {
	tests := []struct {
		name string
		id   ID
		want string
	}{
		{
			name: "case 1",
			id:   3454662929425629184,
			want: "2vsbdk5ri0000",
		},
		{
			name: "case 2",
			id:   3454693972411154785,
			want: "2vsc9rksk00b1",
		},
		{
			name: "case 3",
			id:   3454693972413251600,
			want: "2vsc9rksm000g",
		},
		{
			name: "case 4",
			id:   3454693972413251645,
			want: "2vsc9rksm001t",
		},
		{
			name: "case 5",
			id:   3454693972415348813,
			want: "2vsc9rkso002d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.id.IdStr(); got != tt.want {
				t.Errorf("IdStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
