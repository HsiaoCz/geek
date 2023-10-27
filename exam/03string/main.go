package main

import (
	"fmt"
	"unicode"
)

// 写一个函数判断汉字的字数

func CountHan(str string) (int, string) {
	var hans []rune
	var count int
	for _, char := range str {
		if unicode.Is(unicode.Han, char) {
			count++
			hans = append(hans, char)
		}
	}
	return count, string(hans)
}

func main() {
	str := "hello,你好"
	count, hans := CountHan(str)
	fmt.Println(str, "其中汉字的个数为:", count, "其中的汉字为:", hans)
}
