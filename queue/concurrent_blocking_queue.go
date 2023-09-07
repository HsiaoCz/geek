package queue

import (
	"context"
	"errors"
	"sync"
)

type ConBlockQueue[T any] struct {
	data []T
	lock sync.Mutex
}

func (c *ConBlockQueue[T]) QPush(ctx context.Context, data T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data = append(c.data, data)
	return nil
}

func (c *ConBlockQueue[T]) QPop(ctx context.Context) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.IsEmpty() {
		var data T
		return data, errors.New("空的队列")
	}
	t := c.data[0]
	c.data = c.data[1:]
	return t, nil
}

func (c *ConBlockQueue[T]) IsFull() bool {
	return false
}

func (c *ConBlockQueue[T]) IsEmpty() bool {
	return false
}

func (c *ConBlockQueue[T]) Len() uint64 {
	return 0
}

func New[T any](size int) *ConBlockQueue[T] {
	return &ConBlockQueue[T]{
		data: make([]T, 0, size),
	}
}
