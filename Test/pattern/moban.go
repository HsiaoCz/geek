package pattern

import "fmt"

// 模板方法模式
// 模板方法模式是一种代码复用技术，它提取了类库中公共的行为
// 将公共的行为放在父类中，通过子类来实现不同的行为，它鼓励我们通过继承来实现代码复用

type Beverage interface {
	BoilWater()          //煮开水
	Brew()               // 冲泡
	PourInCup()          // 倒入杯中
	AddThings()          //加料
	WantAddThings() bool // 是否加料
}

// 封装一套流程模板
type template struct {
	b Beverage
}

// 封装固定的模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	// 子类可以重写这个方法来决定是否执行下面的动作
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// 具体的模板子类，制作咖啡
type MakeCoffee struct {
	template // 继承模板
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b 为Beverage，是makeCoffee的接口，这里需要给接口复制，指向具体的子类对象
	// 来触发b全部的接口方法的多态特性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() { fmt.Println("将水煮到100摄氏度") }
func (mc *MakeCoffee) Brew()      { fmt.Println("用水冲咖啡豆...") }
func (mc *MakeCoffee) PourInCup() { fmt.Println("将充好的咖啡到如陶瓷杯中...") }
func (mc *MakeCoffee) AddThings() { fmt.Println("添加牛奶和糖...") }
func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

// 具体的模板子类，制作茶
type MakeTea struct {
	template // 继承模板
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}
func (mt *MakeTea) BoilWater()          { fmt.Println("将水煮到80摄氏度") }
func (mt *MakeTea) Brew()               { fmt.Println("用水冲茶叶...") }
func (mt *MakeTea) PourInCup()          { fmt.Println("将冲好的茶导入茶壶中...") }
func (mt *MakeTea) AddThings()          { fmt.Println("添加柠檬...") }
func (mt *MakeTea) WantAddThings() bool { return false }
