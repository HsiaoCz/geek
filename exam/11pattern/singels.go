package pattern

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例模式的饿汉式
type singeltoon struct{}

var instancee *singeltoon

var initialize uint32

var lock sync.Mutex

func GetInstancee() *singeltoon {
	if atomic.LoadUint32(&initialize) == 1 {
		return instancee
	}
	lock.Lock()
	defer lock.Unlock()
	if initialize == 0 {
		instancee = new(singeltoon)
		atomic.StoreUint32(&initialize, 1)
	}
	return instancee
}

func (s *singeltoon) DoSomething() {
	fmt.Println("单例的某个方法")
}
