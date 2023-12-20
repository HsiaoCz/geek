package pattern

import "fmt"

// 抽象工厂方法模式
// 抽象的水果类
type AbsApple interface {
	ShowApple()
}

type AbsBanana interface {
	ShowBanana()
}

type AbsPear interface {
	ShowPear()
}

// 抽象的工厂类
type AbsFactory interface {
	CreateAbsApple() AbsApple
	CreateAbsBanana() AbsBanana
	CreateAbsPear() AbsPear
}

// 中国的产品族

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() { fmt.Println("this is china apple") }

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() { fmt.Println("this is china banana") }

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() { fmt.Println("this is china pear") }

// 中国的工厂
type ChinaFactory struct{}

func (cf *ChinaFactory) CreateAbsApple() AbsApple {
	return new(ChinaApple)
}
func (cf *ChinaFactory) CreateAbsBanana() AbsBanana {
	return new(ChinaBanana)
}
func (cf *ChinaFactory) CreateAbsPear() AbsPear {
	return new(ChinaPear)
}

// 同理可以实现美国的产品族
