package anymore

import "log"

// 分组
type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	parent      *RouterGroup
	engine      *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRouter(method string, comp string, handler HandleFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRouter(method, pattern, handler)
}
func (g *RouterGroup) GET(pattern string, handler HandleFunc) {
	g.addRouter("GET", pattern, handler)
}
func (g *RouterGroup) POST(pattern string, handler HandleFunc) {
	g.addRouter("POST", pattern, handler)
}
func (g *RouterGroup) PUT(pattern string, handler HandleFunc) {
	g.addRouter("PUT", pattern, handler)
}
func (g *RouterGroup) DELETE(pattern string, handler HandleFunc) {
	g.addRouter("DELETE", pattern, handler)
}
func (g *RouterGroup) HEAD(pattern string, handler HandleFunc) {
	g.addRouter("HEAD", pattern, handler)
}
func (g *RouterGroup) TRANCE(pattern string, handler HandleFunc) {
	g.addRouter("TRANCE", pattern, handler)
}

func (g *RouterGroup) OPTIONS(pattern string, handler HandleFunc) {
	g.addRouter("OPTIONS", pattern, handler)
}
func (g *RouterGroup) PATCH(pattern string, handler HandleFunc) {
	g.addRouter("PATCH", pattern, handler)
}
func (g *RouterGroup) CONNECT(pattern string, handler HandleFunc) {
	g.addRouter("CONNECT", pattern, handler)
}
func (g *RouterGroup) Use(middleware ...HandleFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}
