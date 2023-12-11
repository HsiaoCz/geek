package main

import "fmt"

// [1,2,3,4,5]
func main() {
	a := []int{1, 2, 3, 4, 5}
	s := ON(a)
	fmt.Println(s)
	Mem(a)
	fmt.Println(a)
}

// 这种方式需要重新开辟内存
func ON(a []int) []int {
	b := []int{}
	if len(a) == 0 {
		return a
	}
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}

// 不需要开辟内存的
func Mem(a []int) {
	if len(a) == 0 {
		return
	}
	j := len(a) - 1
	for i := 0; i < len(a); i++ {
		if i == j {
			break
		}
		a[i] = a[i] - a[j]
		a[j] = a[j] + a[i]
		a[i] = a[j] - a[i]
		j = j - 1
	}
}
