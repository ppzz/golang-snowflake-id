package id

import (
	"strconv"
)

type ID int64

func NewId(i int64) ID {
	return ID(i)
}

// Int64 返回一个id的int64值
func (id ID) Int64() int64 {
	return int64(id)
}

// ToInt64 返回一个id的int64值
func (id ID) ToInt64() int64 {
	return id.Int64()
}

// GetTimestamp 从ID中取出时间戳 ms
func (id ID) GetTimestamp() int64 {
	temp := unChaos(uint64(id))
	return int64(temp) >> (serverIdLength + counterLength)
}

// GetServerID 从ID中取出serverID
func (id ID) GetServerID() int64 {
	temp := unChaos(uint64(id))
	return (serverIdSchema & int64(temp)) >> counterLength
}

// GetCounter 从ID中取出序号
func (id ID) GetCounter() int64 {
	temp := unChaos(uint64(id))
	return counterSchema & int64(temp)
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

	for i := 0; i < 13; i++ { // 13 字符串的长度
		v = mask5bit & val
		result = append(result, listIntToByte[v])
		val = val >> 5
	}
	reverse(result)
	return string(result)
}

// IdShortStr 64进制字符串
func (id ID) IdShortStr() string {
	val := id.Int64()
	v := int64(0)
	result := make([]byte, 0)

	for i := 0; i < 11; i++ {
		v = mask6bit & val
		result = append(result, list64Char[v])
		val = val >> 6
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
var listIntToByte2 = "0123456789abcdefghijklmnopqrstuv"
var list64Char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"

// 没有采用通常的32进制字符，而是使用去掉不易辨认的字符后的32个字符编码
// var listIntToByte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

var mask5bit = int64(0x1f)
var mask6bit = int64(0x3f)

// 上面 listIntToByte 的 反向集合
var listByteToInt []int64
var list64CharReverse []int64

func init() {
	maxCharLen := 0
	for _, item := range listIntToByte {
		if int(item) > maxCharLen {
			maxCharLen = int(item)
		}
	}

	listByteToInt = make([]int64, maxCharLen+1)
	for idx, item := range listIntToByte {
		listByteToInt[item] = int64(idx)
	}

	maxCharLen2 := 0
	for _, item := range list64Char {
		if int(item) > maxCharLen2 {
			maxCharLen2 = int(item)
		}
	}
	list64CharReverse = make([]int64, maxCharLen2+1)
	for idx, item := range list64Char {
		list64CharReverse[item] = int64(idx)
	}
}
