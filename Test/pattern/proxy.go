package pattern

import "fmt"

// 代理模式
// 代理模式就是与被代理的对象具有相同的接口的类
// 客户端通过代理与目标类进行交互，代理一般在交互的过程中
// 可以进行一些特别的处理

type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品的真假
}

// 抽象的主题
type Shopping interface {
	Buy(goods *Goods) // 代理的任务
}

// 实现层
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物,买了:", goods.Kind)
}

// 具体的购物主题
type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物,买了:", goods.Kind)
}

// / 设置代理
type OverseasProxy struct {
	shoping Shopping //代理某个主题
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shoping: shopping}
}

func (o *OverseasProxy) Buy(goods *Goods) {
	if o.distinguish(goods) {
		//2. 进行购买
		o.shoping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		o.check(goods)
	}
}

func (o *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if !goods.Fact {
		fmt.Println("发现假货", goods.Kind, ",不应该购买.")
	}
	return goods.Fact
}

func (o *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查,成功带回祖国")
}
