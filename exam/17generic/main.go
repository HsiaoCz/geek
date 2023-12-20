package main

type MySlice[T int | int32 | float32] []T

// T 类型形参 T的具体类型不确定，类似于占位符
// int | int32 | float32 类型约束 T只能接收这几个中的一个
// []里面的为类型的参数列表，MySlice中只有一个T
// 定义的泛型类型名称MySlice[T]

var a MySlice[int]

func main() {
	a = append(a, 1)
}
