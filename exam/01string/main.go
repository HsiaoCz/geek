package main

import "fmt"

// 1、反转字符串
func main() {
	str := "hello,张三"
	astr := FanStr(str)
	fmt.Println("反转之后的字符串:", astr)
}

func FanStr(str string) string {
	// 反转字符串
	// 将字符串转成字节切片，追加到另一个切片
	// 然后将切片转换成字符串
	var astr []rune
	for i := len([]rune(str)) - 1; i >= 0; i-- {
		astr = append(astr, []rune(str)[i])
	}
	return string(astr)
}
