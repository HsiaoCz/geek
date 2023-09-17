package queue

import "sync"

type Mqueue struct {
	lock   sync.Mutex
	data   *node
	length int
	maps   map[int]struct{}
}
