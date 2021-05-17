package id

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"sync/atomic"
	"time"
)

/*
 * int64位的ID生成器
 * 编码规则: 留一位给系统做符号为，剩余63位可以作为ID的有效位
 * 1     + 42        + 11        + 10
 * 符号位 + 时间戳(ms) + Server ID + 自增ID
 * 1  bit: 不使用
 * 42 bit: max: 0x3FFFFFFFFFF == 4,398,046,511,103 note: 最大"2109-05-15 15:35:11.103 +0800 CST"
 * 11 bit: max:         0x7FF ==             2,047 note: 最多2047个
 * 10 bit: max:         0x3FF ==             1,023 note: 最多约1023个，再多就会重复
 */

// 程序的配置
const (
	timestampLength = 42                 // 位数
	serverIdLength  = 11                 // 位数
	counterLength   = 10                 // 位数
	timeStampMax    = 0x3FFFFFFFFFF      // 对应bit位的最大值
	serverIdMax     = 0x7FF              // 对应bit位的最大值
	counterMax      = 0x3FF              // 对应bit位的最大值
	timeStampSchema = timeStampMax << 21 // 最大值在指定位置的bit格式，用来参与位运算
	serverIdSchema  = serverIdMax << 10  // 最大值在指定位置的bit格式，用来参与位运算
	counterSchema   = counterMax         // 最大值在指定位置的bit格式，用来参与位运算
)

const (
	counterStart        = 0 // counter从0开始累加，每次+1之后返回
	counterIncreaseStep = 1 // counter 自增的步长
)

var (
	lastTimeStamp = int64(0) // 上一次生成ID时的 timestamp, ms
	serverID      = int64(0) // 服务器ID
	counter       = int64(0) // 自增计数器
)

// Init 初始化,  如果没有初始化，生成出的ID的 serverID 部分为0
func Init(sid int32) {
	serverID = int64(sid)
}

// Generate 获取一个唯一ID，线程安全，无重复风险
func Generate() ID {
	// counter达到最大值，等待1ms
	if counter == counterMax {
		time.Sleep(time.Millisecond)
	}
	ms := msTimeStamp()

	// 如果当前时间跟上次记录的时间不一样: 重置counter
	if ms != lastTimeStamp {
		lastTimeStamp = ms
		atomic.AddInt64(&counter, counterStart-counter)
	}

	counter := atomic.AddInt64(&counter, counterIncreaseStep)
	return assemble(ms, serverID, counter)
}

// GenerateWithDetail 仅供测试使用，业务逻辑请使用Generate方法
func assemble(timestampMs, serverId, counter int64) ID {
	if timestampMs&timeStampMax != timestampMs {
		panic("assemble id: timestamp is not correct." + strconv.Itoa(int(timestampMs)))
	}
	if serverId&serverIdMax != serverId {
		panic("assemble id: serverID is not correct." + strconv.Itoa(int(timestampMs)))
	}
	if counter&counterMax != counter {
		panic("assemble id: counter is not correct." + strconv.Itoa(int(timestampMs)))
	}
	i := (timestampMs << (serverIdLength + counterLength)) | (serverId << counterLength) | counter
	return ID(i)
}

// FromStr 字符串转换为ID
func FromStr(s string) ID {
	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		logrus.Errorf("FromStr failed: %v", s)
	}
	return ID(i)
}

func msTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
