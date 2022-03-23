package id

import (
	"strconv"
)

type ID int64

// Int64 返回一个id的int64值
func (id ID) Int64() int64 {
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

// IdStr 32进制字符串
func (id ID) IdStr() string {
	val := id.Int64()
	v := int64(0)
	result := make([]byte, 0)

	for i := 0; i < 13; i++ {
		v = mask5bit & val
		result = append(result, listIntToByte[v])
		val = val >> 5
	}
	reverse(result)
	return string(result)
}

// String 重写 Stringer 接口的的 String 方法
func (id ID) String() string {
	return id.IdStr()
}

func reverse(b []byte) {
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
}

var listIntToByte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v'}

// 没有采用通常的32进制字符，而是使用去掉不易辨认的字符后的32个字符编码
// var listIntToByte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

var mask5bit = int64(0x1f)

// 上面 listIntToByte 的 反向集合
var listByteToInt []int64

func init() {
	max := 0
	for _, item := range listIntToByte {
		if int(item) > max {
			max = int(item)
		}
	}

	listByteToInt = make([]int64, max+1)
	for idx, item := range listIntToByte {
		listByteToInt[item] = int64(idx)
	}
}
