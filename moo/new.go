package moo

// 这里使用选项模式
// 在初始化Engine的时候可以添加额外的配置

func New(opts ...HTTPOption) *Engine {
	e := &Engine{
		router: newRouter(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
