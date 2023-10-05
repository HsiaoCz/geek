package lslice

import "errors"

// 二等奖的slcie

type scslice struct {
	name string
	fs   []int
}

var scins = new(scslice)

// 二等奖的单例
func GetScins() *scslice {
	return scins
}

// 这里设置奖品的名称
// 当然奖品可能还有很多别的东西
func (s *scslice) SetName(name string) {
	s.name = name
}

// 获取奖品的名称
func (s *scslice) GetName() string {
	return s.name
}

// 设置中奖的num
func (s *scslice) Setfs(num ...int) {
	s.fs = append(s.fs, num...)
}

// 删除内容
// 删除最后一个值
func (s *scslice) Delfs() error {
	if len(s.fs) == 0 {
		return errors.New("切片为空")
	}
	s.fs = s.fs[:len(s.fs)-1-1]
	return nil
}

// 清空切片
func (s *scslice) Clefs() error {
	if len(s.fs) == 0 {
		return errors.New("切片为空")
	}
	s.fs = []int{}
	return nil
}

// 获取中奖的slcie
// 这里主要是用来遍历取值的
func (s *scslice) Getfs() []int {
	return s.fs
}
