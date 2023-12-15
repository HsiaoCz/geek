package pattern

import "fmt"

// 工厂方法模式
// 简单工厂模式不满足开闭原则
// 工厂方法模式在简单工厂方法模式之上增加了开闭原则

type FFruit interface {
	Show()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateFruit() FFruit
}

// 实现层
type Apples struct{}

func (a *Apples) Show() {
	fmt.Println("this is apple")
}

type Bananas struct{}

func (b *Bananas) Show() {
	fmt.Println("this is banana")
}

type Pears struct{}

func (p *Pears) Show() {
	fmt.Println("this is pear")
}

// 工厂类
type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() FFruit {
	return new(Apples)
}

type BananaFactory struct{}

func (bf *BananaFactory) CreateFruit() FFruit {
	return new(Bananas)
}

type PearFactory struct{}

func (pf *PearFactory) CreateFruit() FFruit {
	return new(Pears)
}
