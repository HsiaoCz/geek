package main

import (
	"fmt"
	"unicode"
)

// 写一个函数，可以判断一个字符串中汉字的字数

func CountHan(str string) int {
	count := 0
	for _, char := range str {
		if unicode.Is(unicode.Han, char) {
			count++
		}
	}
	return count
}

// 这里再进一步
// 不仅统计字符串中汉字的个数
// 将汉字作为一个字符串输出

func CountAndStrHan(str string) (count int, hans string) {
	han := make([]rune, 0)
	for _, char := range str {
		if unicode.Is(unicode.Han, char) {
			count++
			han = append(han, char)
		}
	}
	hans = string(han)
	return
}

func main() {
	str := "hello张三"
	fmt.Println("hello张三中汉字的个数为:", CountHan(str))
	str1 := "hihi李四你好"
	num, hans := CountAndStrHan(str1)
	fmt.Printf("%s:中的汉字数为%d,汉字为:%s\n", str1, num, hans)
}
