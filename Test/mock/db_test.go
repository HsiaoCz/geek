package db

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get()方法是否被调用

	m := NewMockDB(ctrl)
	// Return 返回确定的值
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	// Do Mock 方法被调用时，要执行的操作吗，忽略返回值
	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	// DoAndReturn 可以动态地控制返回值
	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1,but got", v)
	}
}

// 在上面的例子中，当 Get() 的参数为 Tom，则返回 error，这称之为打桩(stub)，有明确的参数和返回值是最简单打桩方式
// 除此之外，检测调用次数、调用顺序，动态设置返回值等方式也经常使用

// m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
// m.EXPECT().Get(gomock.Any()).Return(630, nil)
// m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
// m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

// Eq(value) 表示与 value 等价的值
// Any() 可以用来表示任意的入参。
// Not(value) 用来表示非 value 以外的值。
// Nil() 表示 None 值

// 调用次数
func TestGetFromDB1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")
}

// Times() 断言 Mock 方法被调用的次数。
// MaxTimes() 最大次数。
// MinTimes() 最小次数。
// AnyTimes() 任意次数（包括 0 次）

// 调用顺序
func TestGetFromDB2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}
