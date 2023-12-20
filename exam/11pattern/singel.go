package pattern

import "fmt"

// 单例模式
// 单例模式的懒汉式

type singelton struct{}

var instance = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) DoSomething() {
	fmt.Println("单例的某个方法")
}
