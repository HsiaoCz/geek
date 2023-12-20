package pattern

import "fmt"

// 简单工厂模式
// 简单工厂模式实现了对象创建和使用的分离
// 缺点在于工厂类的职责很重,而且不满足开闭原则

// 抽象
type Fruit interface {
	Show()
}

// 实现层
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("this is banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("this is pear")
}

// 工厂类
type SimpleFactory struct{}

func (s *SimpleFactory) NewFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	}
	if kind == "banana" {
		fruit = new(Banana)
	}
	if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}
