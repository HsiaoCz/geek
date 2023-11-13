package cache

import (
	"container/list"
)

// LRU最近最少使用
// 如果数据最近被访问过，那么将来被访问的概率也会更高
// LRU的实现，使用队列，如果某条记录被访问了，那么将其移到到队尾
// 那么队首元素就是最近最少访问的，将其删掉就可以了

type Cache struct {
	// 允许使用的最大内存
	maxBytes int64
	// 当前已使用的内存
	nbytes int64
	// 双向链表
	ll *list.List
	// map结构用来存储缓存，键为string,值为链表中一个节点的指针
	cache map[string]*list.Element
	// 记录某条记录被移除时的回调函数
	OnEvicted func(key string, value Value)
}

// entry 双向链表节点的数据类型
// 这里仍然记录Key value 这样在淘汰队首元素的时候，直接使用key从Map中删除
type entry struct {
	key   string
	value Value
}

// 这里的值可以是实现了Value的任意类型
// 这个接口仅有一个方法
// 返回值本身的大小

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找功能
// 第一步需要从字典中找到对应的双向链表的节点
// 第二步，将该节点移动到队尾

func (c *Cache) Get(key string) (value Value, ok bool) {
	// 判断传递来的key是否在map中
	// 是 则将值移动到队尾
	// 并将查询到的值返回，并返回是否查询到值
	if ele, ok := c.cache[key]; ok {
		// MoveToFront()
		// 将值移动到队尾
		// 但是双向链表，队首和队尾是相对的
		// 这里以front为队尾
		c.ll.MoveToFront(ele)
		// 对值进行断言
		// 返回查询到的值和bool
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除功能
// 删除就是淘汰策略
// 淘汰最近最少没有访问的，即队首元素

func (c *Cache) RemoveOldest() {
	// 先在这里把链表里面的队首元素返回
	ele := c.ll.Back()
	// 如果元素不为空
	if ele != nil {
		// 移除元素
		c.ll.Remove(ele)
		// 对元素进行断言
		kv := ele.Value.(*entry)
        // 删除map里的对应的键值对
		// 这里就体现出entry 结构体里面放key的好处了
		delete(c.cache, kv.key)
		// 重新计算当前内存的使用状况
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		// 假如删除元素时的回调函数不为空
		if c.OnEvicted != nil {
			// 则执行该回调函数
			c.OnEvicted(kv.key, kv.value)
		}
	}
}
