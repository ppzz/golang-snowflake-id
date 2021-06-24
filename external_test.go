package id

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate_A(t *testing.T) {
	ids := make([]ID, 100000)
	for i := range ids {
		ids[i] = Generate()
	}
	lastC := int64(0)
	lastTS := int64(0)

	t.Run("case1: 验证生成的ID是否递增", func(t *testing.T) {
		for _, id := range ids {
			c := id.GetCounter()
			ms := id.GetTimestamp()
			if lastC == 1023 {
				lastC = 0
			}
			if lastTS != ms {
				lastC = 0
				lastTS = ms
			}
			if lastC+1 != c {
				format := "counter 没有递增() = lastCounter: %d, c: %d, lastTS: %d, ts: %d, server: %d"
				t.Errorf(format, lastC, c, lastTS, ms, id.GetServerID())
			}
			lastC = c
		}
	})
}

func TestFromStr(t *testing.T) {
	type args struct {
		ts  int64
		sid int64
		c   int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				ts:  msTimeStamp(),
				sid: 1000,
				c:   123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := assemble(tt.args.ts, tt.args.sid, tt.args.c)
			got := FromStr(id.ToStr())
			fmt.Println(id.ToStr())
			if got != id {
				t.Errorf("FromStr() = %v, want %v", got, id)
			}
		})
	}
}

func Test_durationToNextMillisecond(t *testing.T) {
	tests := []struct {
		name string
		want time.Duration
	}{
		{
			name: "case1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toNextMillisecond()
			fmt.Println(time.Now().Nanosecond(), got)
			if got != tt.want {
				t.Errorf("toNextMillisecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
