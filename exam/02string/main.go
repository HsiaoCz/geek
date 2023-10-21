package main

import "fmt"

// 判断一个字符串是不是回文字符串

func main() {
	str := "我是不是我"
	if ok := HStr(str); ok {
		fmt.Printf("\"%s\":是回文字符串\n", str)
	} else {
		fmt.Printf("\"%s\":不是回文字符串\n", str)
	}
}

func HStr(str string) bool {
	var astr []rune
	for i := len([]rune(str)) - 1; i >= 0; i-- {
		astr = append(astr, []rune(str)[i])
	}
	if mstr := string(astr); mstr == str {
		return true
	}
	return false
}
