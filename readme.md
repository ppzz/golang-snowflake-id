# Snowflake ID Generator

## desc

* ID base on int64(基于 int64 的ID结构)
* concurrency safe(并发安全)
* max 1023 ID per millisecond(每毫秒最多1023个ID)
* will sleep to next millisecond when reach limit in one millisecond(某一毫秒内达到最大值会sleep到下一毫秒)

## design

ID is technically an int64:

* consist of 4 parts: (reserved sign, millisecond timestamp, serverId, auto increase counter)
  * reserved sign( 预留的符号位 ): 1 bit,
  * millisecond timestamp( 毫秒时间戳 ): 42 bit
  * serverId(self defined,can be any val),( 自定义ID ): 11 bit
  * auto increase counter( 自增计数器 ): 10 bit

|desc|bit| max| max|note|
|---|---|---|---|---|
|reserved sign| 1|0x1|1|not use|
|millisecond timestamp| 42|0x3FFFFFFFFFF|  4,398,046,511,103| max to "2109-05-15 15:35:11.103 +0800 CST"|
|serverId| 11|0x7FF|  2,047 |-|
|auto increase counter| 10|0x3FF|  1,023 |-|

## usage

### install package:

```bash
go get github.com/ppzz/golang-snowflake-id
```

### use:

import:

```go
import id "github.com/ppzz/golang-snowflake-id"
```

Set ServerID(optional)

default 0 if not set

```go
sId := 1001
id.Init(sId)
```

Set DisableLog(optional)

by default: will print a log before thread sleeping(counter reach max val)

```go
id.DisableLog()
```

Generate New ID:

```go
newId := id.Generate()
```

Get Info from an ID object:

```go
fmt.Println(newId.ToStr())
fmt.Println(newId.ToInt64())
fmt.Println(newId.GetCounter())
fmt.Println(newId.GetTimestamp())
fmt.Println(newId.GetServerID())
```
