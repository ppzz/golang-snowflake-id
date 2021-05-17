package id

import "strconv"

type ID int64

func (id ID) Hello() {

	return
}

func (id ID) ToInt64() int64 {
	return int64(id)
}

func (id ID) ToInt() int {
	return int(id)
}

// GetTimestamp 从ID中取出时间戳 ms
func (id ID) GetTimestamp() int64 {
	return int64(id) >> (serverIdLength + counterLength)
}

// GetServerID 从ID中取出serverID
func (id ID) GetServerID() int64 {
	return (serverIdSchema & int64(id)) >> counterLength
}

// GetCounter 从ID中取出序号
func (id ID) GetCounter() int64 {
	return counterSchema & int64(id)
}

// ToStr 将ID字符串化
func (id ID) ToStr() string {
	return strconv.FormatInt(int64(id), 16)
}
