package santino

import "net/http"

// 我们开始实现一个自己的框架
// 第一步，实现一个自己的Handler

type Santino struct{}

// http.ListenAndServe(addr,handler)
// 第二个参数handler本质是一个接口
// 这个接口又一个方法
// 就是ServeHTTP(http.ResponseWriter,*http.Request)
// 我们实现这个接口，就有了自己的handler
func (s *Santino) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// 这里的new方法类似于gin.New()返回一个默认的Santino
func New() *Santino {
	return &Santino{}
}
