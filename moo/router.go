package moo

import (
	"fmt"
	"strings"
)

// 路由树
// 一个方法一个路由树
type router struct {
	trees map[string]*node
}

func newRouter() *router {
	return &router{
		trees: make(map[string]*node),
	}
}

func (r *router) addRouter(method string, pattern string, handlefunc Handlefunc) {
	// 打印一下注册的路由
	fmt.Printf("add router %s - %s\n", method, pattern)
	if pattern == "" {
		panic("web:路由不能为空")
	}
	// 获取根节点
	root, ok := r.trees[method]
	if !ok {
		// 根节点不存在
		// 1.创建根节点
		// 2.把根节点放到trees里面
		root = &node{
			part: "/",
		}
		r.trees[method] = root
	}
	// TODO 如果是以根路由怎么办?

	if pattern == "/" {
		root.handlefunc = handlefunc
		return
	}

	if !strings.HasPrefix(pattern, "/") {
		panic("web:路由必须以 / 开头")
	}
	if strings.HasSuffix(pattern, "/") {
		panic("web:路由不能以 / 结尾")
	}

	// 切割pattern
	// /user/login => ["","user","login"]
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		if part == "" {
			panic("web:路由不能出现连续的 / ")
		}
		root = root.addNode(part)
	}
	root.handlefunc = handlefunc
}

func (r *router) getRouter(method string, pattern string) (*node, map[string]string, bool) {
	params := make(map[string]string)
	if pattern == "" {
		return nil, params, false
	}
	// TODO  / 这种路由怎么办？
	// 获取根节点

	root, ok := r.trees[method]
	if !ok {
		return nil, params, false
	}
	if pattern == "/" {
		return root, params, true
	}
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		if part == "" {
			return nil, params, false
		}
		root = root.getNode(part)
		if root == nil {
			return nil, params, false
		}
		if strings.HasPrefix(root.part, ":") {
			params[root.part[1:]] = part
		}
		// /stufy/:course/action
	}
	return root, params, root.handlefunc != nil
}
