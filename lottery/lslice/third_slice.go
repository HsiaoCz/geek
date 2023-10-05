package lslice

import "errors"

// 一等奖的slcie

type tslice struct {
	name string
	fs   []int
}

var tins = new(tslice)

func GetTins() *tslice {
	return tins
}

// 这里设置奖品的名称
// 当然奖品可能还有很多别的东西
func (t *tslice) SetName(name string) {
	t.name = name
}

// 获取奖品的名称
func (t *tslice) GetName() string {
	return t.name
}

// 设置中奖的num
func (t *tslice) Setfs(num ...int) {
	t.fs = append(t.fs, num...)
}

// 删除内容
// 删除最后一个值
func (t *tslice) Delfs() error {
	if len(t.fs) == 0 {
		return errors.New("切片为空")
	}
	t.fs = t.fs[:len(t.fs)-1-1]
	return nil
}

// 清空切片
func (t *tslice) Clefs() error {
	if len(t.fs) == 0 {
		return errors.New("切片为空")
	}
	t.fs = []int{}
	return nil
}

// 获取中奖的slcie
// 这里主要是用来遍历取值的
func (t *tslice) Getfs() []int {
	return t.fs
}
