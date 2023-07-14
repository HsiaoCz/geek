package context

import (
	"testing"
)

func TestContext(t *testing.T) {
	// 一般是链路起点，或者调用的起点
	// ctx := context.Background()
	// 不确定context该用啥的时候用TODO
	// ctx1 := context.TODO()

	// context.WithValue
	// ctx = context.WithValue(ctx, "my-key", "my-value")
	// ctx, cancel := context.WithCancel(ctx)
	// 用完ctx 再去调用
	// cancel()
}

func TestContextParent(t *testing.T) {
	// ctx := context.Background()

	// parent := context.WithValue(ctx, "mykey", "myValue")

	// child := context.WithValue(parent, "mykey", "my new value")

	// t.Log("parent mykey:", parent.Value("mykey"))
	// t.Log("child mykey:", child.Value("mykey"))
}
