package id

import (
	"strconv"
	"time"
)

// msTimeStamp 毫秒为单位的时间戳
func msTimeStamp() uint64 {
	return uint64(time.Now().UnixNano() / 1e6)
}

// GenerateWithDetail 仅供测试使用，业务逻辑请使用Generate方法
func assemble(timestampMs, serverId, counter uint64) ID {
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
	return NewId(int64(chaos(i)))
}

// toNextMillisecond
func toNextMillisecond() time.Duration {
	microSecondTimeStamp := time.Now().Nanosecond() / 1000 // 当前的微秒数的时间戳
	millisecond := microSecondTimeStamp % 1000             // 当前的微秒数
	needSleep := 1000 - millisecond + 1                    // 需要等待的微秒数
	return time.Microsecond * time.Duration(needSleep)
}

func unChaos(ii uint64) uint64 {
	f := getSubBits(ii, 0, 1)
	t1 := getSubBits(ii, 1, 11)  // 11 bit
	s1 := getSubBits(ii, 12, 8)  // 8 bit
	t2 := getSubBits(ii, 20, 16) // 16 bit
	c1 := getSubBits(ii, 36, 8)  // 8 bit
	t3 := getSubBits(ii, 44, 12) // 12 bit
	s2 := getSubBits(ii, 56, 3)  // 3 bit
	t4 := getSubBits(ii, 59, 3)  // 3 bit
	c2 := getSubBits(ii, 62, 2)  // 2 bit

	val := uint64(0)
	val = setBits(val, f, 0, 1)    // 1 bit
	val = setBits(val, t1, 1, 11)  // 11 bit
	val = setBits(val, t2, 12, 16) // 16 bit
	val = setBits(val, t3, 28, 12) // 12 bit
	val = setBits(val, t4, 40, 3)  // 3 bit
	val = setBits(val, s1, 43, 8)  // 8 bit
	val = setBits(val, s2, 51, 3)  // 3 bit
	val = setBits(val, c1, 54, 8)  // 8 bit
	val = setBits(val, c2, 62, 2)  // 2 bit
	return val
}

func chaos(ii uint64) uint64 {
	f := getSubBits(ii, 0, 1)
	t1 := getSubBits(ii, 1, 11)  // 11 bit
	t2 := getSubBits(ii, 12, 16) // 16 bit
	t3 := getSubBits(ii, 28, 12) // 12 bit
	t4 := getSubBits(ii, 40, 3)  // 3 bit
	s1 := getSubBits(ii, 43, 8)  // 8 bit
	s2 := getSubBits(ii, 51, 3)  // 3 bit
	c1 := getSubBits(ii, 54, 8)  // 8 bit
	c2 := getSubBits(ii, 62, 2)  // 2 bit

	val := uint64(0)
	val = setBits(val, f, 0, 1)    // 1 bit
	val = setBits(val, t1, 1, 11)  // 11 bit
	val = setBits(val, s1, 12, 8)  // 8 bit
	val = setBits(val, t2, 20, 16) // 16 bit
	val = setBits(val, c1, 36, 8)  // 8 bit
	val = setBits(val, t3, 44, 12) // 12 bit
	val = setBits(val, s2, 56, 3)  // 3 bit
	val = setBits(val, t4, 59, 3)  // 3 bit
	val = setBits(val, c2, 62, 2)  // 2 bit

	return val
}

// setBits 设置int64的某一段bit
func setBits(val uint64, use uint64, pos int, bits int) uint64 {
	mask := GenerateBitsMask(bits)
	move := 64 - pos - bits

	// 从 use 中取出需要的部分, 并移动到目标位置
	newUse := use & mask
	newUse = newUse << move

	// 清除原来的位置
	clearMask := mask << move
	clearMask = ^clearMask // 取反, 目标位置将变为0,其他是1
	val = val & clearMask  // 做与操作, 则 目标位置全部变为0, 其他位置不变

	return val | newUse // 做或操作, 将两个值合并
}

// getSubBits 取 int64 的某一段bit
func getSubBits(val uint64, pos int, bits int) uint64 {
	// 10 位开始, 取5个bit
	mask := GenerateBitsMask(bits)
	move := 64 - pos - bits
	return val >> move & mask
}

// GenerateBitsMask 输入一个数字, 返回n个1组成的二进制数字的int64值
func GenerateBitsMask(n int) uint64 {
	var result uint64 = 0
	for i := 0; i < n; i++ {
		result |= 1 << i
	}
	return result
}
