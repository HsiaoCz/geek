package lslice

import "errors"

// 一等奖的slcie

type fslice struct {
	name string
	fs   []int
}

var fins = new(fslice)

func GetFins() *fslice {
	return fins
}

// 这里设置奖品的名称
// 当然奖品可能还有很多别的东西
func (f *fslice) SetName(name string) {
	f.name = name
}

// 获取奖品的名称
func (f *fslice) GetName() string {
	return f.name
}

// 设置中奖的num
func (f *fslice) Setfs(num ...int) {
	f.fs = append(f.fs, num...)
}

// 删除内容
// 删除最后一个值
func (f *fslice) Delfs() error {
	if len(f.fs) == 0 {
		return errors.New("切片为空")
	}
	f.fs = f.fs[:len(f.fs)-1-1]
	return nil
}

// 清空切片
func (f *fslice) Clefs() error {
	if len(f.fs) == 0 {
		return errors.New("切片为空")
	}
	f.fs = []int{}
	return nil
}

// 获取中奖的slcie
// 这里主要是用来遍历取值的
func (f *fslice) Getfs() []int {
	return f.fs
}
