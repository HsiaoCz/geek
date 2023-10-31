package pattern

import "fmt"

// 装饰器模式
// 动态的给一个对象添加一些额外的职责

// 抽象的手机
type Phone interface {
	Show()
}

// 装饰器的基础类
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 具体的手机
type XiaoMi struct{}

func (x *XiaoMi) Show() { fmt.Println("小米手机") }

type HuaWei struct{}

func (h *HuaWei) Show() { fmt.Println("华为手机") }

// 具体的装饰器
type MoDecoator struct {
	Decorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecoator{Decorator{phone: phone}}
}

func (md *MoDecoator) Show() {
	md.phone.Show()
	fmt.Println("戴上了手机壳的手机...")
}

type KeDecorator struct {
	Decorator
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

func (ke *KeDecorator) Show() {
	ke.phone.Show()
	fmt.Println("戴上了手机壳的手机")
}
