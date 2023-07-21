package queue

import "context"

type Queue[T any] interface {
	// queue的入队和出队的方法
	// 入队的方法
	// 这里的context主要用于考虑超时时间
	QPush(ctx context.Context, data T) error
	// 出队的方法
	QPop(ctx context.Context) (T, error)
	// 队列是不是满了
	IsFull() bool
	// 队列是不是为空
	IsEmpty() bool
	// 获取队列的长度
	Len() uint64
}
