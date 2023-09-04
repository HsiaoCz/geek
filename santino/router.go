package santino

// 路由树
// 这里每个请求方法构造一个路由树
// 这里的树基于前缀树，由于相同父节点的节点具有相同的前缀

type router struct {
	trees map[string]*node
}
