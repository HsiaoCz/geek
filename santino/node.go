package santino

// 这里每个node代表树中的一个节点

type node struct {
	// part 代表当前节点的的参数
	// 这里思考一个问题
	// /user/register
	// /user/login
	// 这个应该怎么存储?
	// / 作为根节点
	// user 作为子节点
	// register login 分别作为子节点
	// 这里还需要考虑一个问题？
	// /也作为存储的一部分吗？
	// 放在前面还是后面？
	part string
}
