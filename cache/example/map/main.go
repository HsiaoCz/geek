package main

import (
	"fmt"
	"sync"
	"time"
)

// go的map是不支持并发的
var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
