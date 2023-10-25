package moo

import "strings"

type node struct {
	part     string
	children map[string]*node
	// 当前节点上存储的视图函数
	handlefunc    Handlefunc
	paramChildren *node
}

func (n *node) addNode(part string) *node {

	if strings.HasPrefix(part, ":") && n.paramChildren == nil {
		n.paramChildren = &node{part: part}
		return n.paramChildren
	}

	// 判断当前节点有没有children属性，就是是不是nil
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[part]
	if !ok {
		child = &node{
			part: part,
		}
		n.children[part] = child
	}
	return child
}

func (n *node) getNode(part string) *node {
	// n的children属性都不存在
	if n.children == nil {
		return nil
	}
	// 正常思路：先到静态路由中找
	child, ok := n.children[part]
	if !ok {
		// 到了这里说明没找到
		// 没找到，说明没有匹配到静态路由
		// 如果动态路由上有值 则返回动态路由
		if n.paramChildren != nil {
			return n.paramChildren
		}
		return nil
	}
	return child
}
