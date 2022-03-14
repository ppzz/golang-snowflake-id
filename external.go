package id

import (
	"errors"
	"log"
	"strconv"
	"sync"
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
	timestampMax    = 0x3FFFFFFFFFF      // 对应bit位的最大值
	serverIdMax     = 0x7FF              // 对应bit位的最大值
	counterMax      = 0x3FF              // 对应bit位的最大值
	timestampSchema = timestampMax << 21 // 最大值在指定位置的bit格式，用来参与位运算
	serverIdSchema  = serverIdMax << 10  // 最大值在指定位置的bit格式，用来参与位运算
	counterSchema   = counterMax         // 最大值在指定位置的bit格式，用来参与位运算
)

const (
	counterStart        = 0 // counter从0开始累加，每次+1之后返回
	counterIncreaseStep = 1 // counter 自增的步长
)

var (
	locker        sync.Mutex
	once          sync.Once
	lastTimestamp = int64(0) // 上一次生成ID时的 timestamp, ms
	serverID      = int64(0) // 服务器ID
	counter       = int64(0) // 自增计数器
	isLogEnable   = true
)

// Init 初始化,  如果没有初始化，生成出的ID的 serverID 部分为0
func Init(sid int32) {
	once.Do(func() {
		serverID = int64(sid)
	})
}

// DisableLog 取消日志
func DisableLog() {
	once.Do(func() {
		isLogEnable = false
	})
}

// Generate 获取一个唯一ID，线程安全，无重复风险
func Generate() ID {
	locker.Lock()
	defer locker.Unlock()

	// case1: 如果当前时间跟上次记录的时间不一样: 重置timestamp&counter, 并返回新id
	ms := msTimeStamp()
	if ms != lastTimestamp {
		lastTimestamp = ms
		counter = counterStart
		return assemble(lastTimestamp, serverID, counter)
	}

	// case2: 当前的时间跟上次时间一致
	// case2.1: counter 没有达到最大值
	if counter < counterMax {
		counter += counterIncreaseStep
		return assemble(lastTimestamp, serverID, counter)
	}

	// case2.2: counter 已经达到最大值，sleep一段时间，重置counter&timestamp
	needSleep := toNextMillisecond()
	if isLogEnable {
		log.Println("id counter reach limit, thread will sleep, counter:", counter, "timestamp:", lastTimestamp, "sleep:", needSleep)
	}
	time.Sleep(needSleep)
	lastTimestamp = msTimeStamp()
	counter = counterStart
	return assemble(lastTimestamp, serverID, counter)
}

// New 获取一个唯一ID，线程安全，无重复风险
func New() ID {
	return Generate()
}

// FromHexStr 16进制字符串转换为ID
func FromHexStr(s string) ID {
	base := 16
	return parse(s, base)
}

// FromDecStr 10进制字符串转换为ID
func FromDecStr(s string) ID {
	base := 10
	return parse(s, base)
}

func parse(s string, base int) ID {
	i, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		panic(errors.New("input str can not parse to ID"))
	}
	return ID(i)
}

// FromString 字符串转换为ID
func FromString(s string) ID {
	return FromHexStr(s)
}

// FromInt64 从int64转化为ID
func FromInt64(i int64) ID {
	return ID(i)
}
