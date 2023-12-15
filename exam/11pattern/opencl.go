package pattern

import "fmt"

// 开闭原则
// 对修改关闭，对扩展开放
// 抽象的银行职员类
type Banker interface {
	Dobuz()
}

// 具体的银行职员
type TransBanker struct{}

func (t *TransBanker) Dobuz() {
	fmt.Println("执行转账操作")
}

// 存款的银行职员类
type SaveBanker struct{}

func (s *SaveBanker) Dobuz() {
	fmt.Println("执行存款操作")
}

// 股票业务的银行职员
type StackBanker struct{}

func (s *StackBanker) Dobuz() {
	fmt.Println("执行股票业务的银行职员")
}
