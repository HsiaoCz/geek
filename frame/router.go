package frame

type Router struct{}

func (r *Router) addRouter(pattern string, handler Handlefunc) {}
func (r *Router) getRouter(pattern string, handler Handlefunc) {}
