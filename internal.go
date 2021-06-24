package id

import (
	"strconv"
	"time"
)

// msTimeStamp 毫秒为单位的时间戳
func msTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// GenerateWithDetail 仅供测试使用，业务逻辑请使用Generate方法
func assemble(timestampMs, serverId, counter int64) ID {
	if timestampMs&timestampMax != timestampMs {
		panic("assemble id: timestamp is not correct." + strconv.Itoa(int(timestampMs)))
	}
	if serverId&serverIdMax != serverId {
		panic("assemble id: serverID is not correct." + strconv.Itoa(int(timestampMs)))
	}
	if counter&counterMax != counter {

		panic("assemble id: counter is not correct.ts:" + strconv.Itoa(int(timestampMs)) + " counter:" + strconv.Itoa(int(counter)))
	}
	i := (timestampMs << (serverIdLength + counterLength)) | (serverId << counterLength) | counter
	return ID(i)
}

// toNextMillisecond
func toNextMillisecond() time.Duration {
	microSecondTimeStamp := time.Now().Nanosecond() / 1000 // 当前的微秒数的时间戳
	millisecond := microSecondTimeStamp % 1000             // 当前的微秒数
	needSleep := 1000 - millisecond + 1                    // 需要等待的微秒数
	return time.Microsecond * time.Duration(needSleep)
}
