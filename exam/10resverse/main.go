package main

import "fmt"

// 反转[1,2,3,4,5]
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(ON(a))
	Mem(a)
	fmt.Println(a)
}

// O(n)
func ON(a []int) []int {
	var b []int
	if len(a) == 0 {
		return a
	}
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}

// mem
func Mem(a []int) {
	if len(a) == 0 {
		return
	}
	j := 0
	for i := len(a) - 1; i >= 0; i-- {
		if i == j {
			break
		}
		a[i] = a[i] - a[j]
		a[j] = a[i] + a[j]
		a[i] = a[j] - a[i]
		j = j + 1
	}
}
