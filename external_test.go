package id

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate_A(t *testing.T) {
	// 产生 100 * 1000 个id
	count := 100 * 000
	ids := make([]ID, count)
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
			got := FromHexStr(id.HexStr())
			fmt.Println(id.HexStr())
			if got != id {
				t.Errorf("FromHexStr() = %v, id %v", got, id)
			}
		})
	}
}

func Test_durationToNextMillisecond(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toNextMillisecond()
			du := time.Duration((time.Now().Nanosecond()/1000)%1000) * time.Microsecond
			// 1001 是表示休眠时间+ 当前时间应该超过一毫秒
			// 1002 表示极端情况下， 上一行里取到的时间可能跟 toNextMillisecond 的时间不在一微秒，所以可能是1002
			if got+du != time.Microsecond*1001 && got+du != time.Microsecond*1002 {
				t.Errorf("toNextMillisecond() = %v, now %v", got, du)
			}
		})
	}
}

func TestFromIdStr(t *testing.T) {
	tests := []struct {
		name  string
		idStr string
		want  ID
	}{
		{
			name:  "case 1",
			idStr: "2vsbdk5ri0000",
			want:  3454662929425629184,
		},
		{
			name:  "case 2",
			idStr: "2vsc9rksk00b1",
			want:  3454693972411154785,
		},
		{
			name:  "case 3",
			idStr: "2vsc9rksm000g",
			want:  3454693972413251600,
		},
		{
			name:  "case 4",
			idStr: "2vsc9rksm001t",
			want:  3454693972413251645,
		},
		{
			name:  "case 5",
			idStr: "2vsc9rkso002d",
			want:  3454693972415348813,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromIdStr(tt.idStr); got != tt.want {
				t.Errorf("FromIdStr() = %v, id %v", got, tt.want)
			}
		})
	}
}
