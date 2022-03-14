package main

import (
	"fmt"
	id "github.com/ppzz/golang-snowflake-id"
	"sort"
)

func main() {
	const goroutineCount = 100
	const idCount = 1000 * 10000
	ch := make(chan id.ID, idCount)

	id.DisableLog()

	for i := 0; i < goroutineCount; i++ {
		go genId(ch, idCount/goroutineCount)
	}

	var list []int
	for {
		i := <-ch
		list = append(list, int(i.ToInt64()))
		if len(list) >= idCount {
			break
		}
	}
	sort.Ints(list)
	uniqList := uniq(list)
	fmt.Println("created ids count:", idCount,
		"\nwith goroutine count:", goroutineCount,
		"\ngenerated-id-list-count:", len(list),
		"\nid-list-count(after-uniq):", len(uniqList))
	close(ch)
}

func uniq(list []int) []int {
	l := len(list)
	if l == 0 {
		return list
	}

	sortedIdx := 0
	for i := 1; i < l; i++ {
		if list[sortedIdx] == list[i] {
			continue
		}
		sortedIdx++
		if sortedIdx < i {
			list[sortedIdx], list[i] = list[i], list[sortedIdx]
		}
	}
	return list[:sortedIdx+1]
}

func genId(ch chan id.ID, count int) {
	for i := 0; i < count; i++ {
		t := id.Generate()
		ch <- t
	}
}
