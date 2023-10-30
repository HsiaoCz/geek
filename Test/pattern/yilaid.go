package pattern

import "fmt"

// 依赖倒转原则
// 依赖于抽象而不是依赖于具体的类

// 抽象的司机
type Driver interface {
	Drive(car Car)
}

type Car interface {
	Run()
}

// 司机类
type Zhangs struct{}

func (z *Zhangs) Drive(car Car) {
	fmt.Println("张三在开车")
	car.Run()
}

// 司机类
type Lisi struct{}

func (l *Lisi) Drive(car Car) {
	fmt.Println("李四在开车")
	car.Run()
}

// 汽车类
type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("奔驰is running...")
}

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("宝马is running...")
}
