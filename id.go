package id

import "strconv"

type ID int64

func (id ID) ToInt64() int64 {
	return int64(id)
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

// HexStr 将ID字符串化（16进制形式）
func (id ID) HexStr() string {
	return strconv.FormatInt(int64(id), 16)
}

// DecStr ID字符串化（十进制形式）
func (id ID) DecStr() string {
	return strconv.FormatInt(int64(id), 16)
}

// String 重写 Stringer 接口的的 String 方法
func (id ID) String() string {
	return id.HexStr()
}

// Int64 返回一个id的int64值
func (id ID) Int64() int64 {
	return int64(id)
}
