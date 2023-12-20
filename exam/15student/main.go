package main

import (
	"fmt"
	"sync"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade int
}

type StuList struct {
	count    int
	l        []*Student
	MaxGrade int
	lock     sync.Mutex
}

func NewStuList() *StuList {
	return &StuList{
		count:    0,
		MaxGrade: 0,
		lock:     sync.Mutex{},
	}
}

func (s *StuList) CreateStudent(stu *Student) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.MaxGrade <= stu.Grade {
		s.MaxGrade = stu.Grade
	}
	s.l = append(s.l, stu)
}

func (s *StuList) RemoveStudent(id int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for index, value := range s.l {
		if value.ID == id {
			s.l = append(s.l[:index], s.l[index+1:]...)
		}
	}
}

func (s *StuList) ListStudent() {
	for _, value := range s.l {
		fmt.Println(value)
	}
}

func (s *StuList) Modefy(id int, stu *Student) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for index, value := range s.l {
		if value.ID == id {
			if s.MaxGrade <= stu.Grade {
				s.MaxGrade = stu.Grade
			}
			s.l[index] = stu
		}
	}
}

func (s *StuList) CountStudent() int {
	return len(s.l)
}

func main() {

}
