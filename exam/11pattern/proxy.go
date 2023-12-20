package pattern

import "fmt"

// proxy 代理模式
// 代理模式就是与被代理的对象具有相同的接口的类
// 客户端通过代理与目标类进行交互，代理一般在交互的过程中
// 可以进行一些特别的处理

type Goods struct {
	Kind string
	Fact bool
}

// 抽象的主题
type Shopping interface {
	Buy(goods *Goods)
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

// 设置代理
type OverseasProxy struct {
	shopping Shopping // 代理某个主题
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping: shopping}
}

func (o *OverseasProxy) Buy(goods *Goods) {
	if o.distinguish(goods) {
		o.shopping.Buy(goods)
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
