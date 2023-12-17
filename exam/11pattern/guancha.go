package pattern

import "fmt"

// 观察者模式
// 观察者模式用于建立一种对象之间的依赖关系，一个对象发生改变将通知其他对象
// 其他对象将作出相应的反应

// 抽象的观察者
type Listener interface {
	OnTeacherComming() // 观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层
// 观察者学生
type StuZhang struct {
	BadThing string
}

func (s *StuZhang) OnTeacherComming() { fmt.Println("张三 停止了...", s.BadThing) }
func (s *StuZhang) DoBadThing()       { fmt.Println("张三 正在", s.BadThing) }

type StuZhao4 struct{ Badthing string }

func (s *StuZhao4) OnTeacherComming() { fmt.Println("赵4 停止 ", s.Badthing) }

func (s *StuZhao4) DoBadthing() { fmt.Println("赵4 正在", s.Badthing) }

type StuWang5 struct{ Badthing string }

func (s *StuWang5) OnTeacherComming() { fmt.Println("王5 停止 ", s.Badthing) }

func (s *StuWang5) DoBadthing() { fmt.Println("王5 正在", s.Badthing) }

// 通知者班长
type ClassMonitor struct {
	listenerList []Listener
}

func (m *ClassMonitor) AddListener(listener ...Listener) {
	m.listenerList = append(m.listenerList, listener...)
}

func (m *ClassMonitor) RemoveListener(listener ...Listener) {
	for _, l := range listener {
		for j, lis := range m.listenerList {
			if l == lis {
				m.listenerList = append(m.listenerList[:j], m.listenerList[j+1:]...)
				break
			}
		}
	}
}

func (m *ClassMonitor) Notifier() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComming()
	}
}
