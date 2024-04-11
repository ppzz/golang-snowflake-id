package main

import (
	"fmt"
	id "github.com/ppzz/golang-snowflake-id"
	strings "strings"
	"time"
)

func main() {
	id.Init(127) // max: 0x7FF == 2047
	id.New()
	id.New()
	id.New()
	id.New()
	newId := id.New()
	fmt.Println("id(int64):", newId.Int64())
	fmt.Println("id(str):  ", newId.String())
	fmt.Println("id(bin):  ", insertSpacesRightToLeft(fill(fmt.Sprintf("%b", newId.Int64()), 64), 4))
	fmt.Println("serverId: ", newId.GetServerID())
	fmt.Println("counter:  ", newId.GetCounter())
	fmt.Println("timestamp:", newId.GetTimestamp(), time.UnixMilli(newId.GetTimestamp()).Format("2006-01-02 15:04:05.000"))
}

func getDesc(name string, val uint64, count int) string {
	i2Str := fmt.Sprintf("%b", val)
	i2Str = fill(i2Str, count)
	i2Str2 := insertSpacesRightToLeft(i2Str, 4)
	return fmt.Sprintf("%10s: %s", name, i2Str2)
}

// insertSpacesRightToLeft 从右往左每隔n个字符插入一个空格
func insertSpacesRightToLeft(s string, n int) string {
	if n <= 0 {
		return s
	}

	result := "" // 使用字符串直接拼接，适用于短字符串操作
	length := len(s)

	// 从字符串末尾开始，每隔n个字符插入一个空格
	for i := length; i > 0; i -= n {
		start := i - n
		if start < 0 {
			start = 0
		}
		// 需要在前面插入空格（除了第一次循环）
		if len(result) > 0 {
			result = s[start:i] + " " + result
		} else {
			result = s[start:i]
		}
	}

	return result
}

func fill(str string, count int) string {
	paddingSize := count - len(str)
	if paddingSize > 0 {
		return strings.Repeat("-", paddingSize) + str
	}
	return str
}
