package moo

func New() *Engine {
	return &Engine{
		router: make(map[string]Handlefunc),
	}
}
