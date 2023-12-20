package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello you you you"
	for key, value := range CountStr(str) {
		fmt.Printf("单词: %s 出现了: %d 次\n", key, value)
	}
	Qimi()
}

func CountStr(str string) map[string]int {
	cstr := make(map[string]int)
	if len(str) == 0 {
		return nil
	}
	for _, s := range strings.Split(str, " ") {
		cstr[s] += 1
	}
	return cstr
}

func Qimi() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
