package main

import "fmt"

type Student struct {
	ID    int
	Age   int
	Grade int
	Name  string
}

var StuList []Student

func AppendStu(stu Student) {
	StuList = append(StuList, stu)
}

func (s *Student) Create(id int, age int, grade int, name string) {
	AppendStu(Student{ID: id, Name: name, Age: age, Grade: grade})
}

func (s *Student) List() {
	for _, v := range StuList {
		fmt.Println(v)
	}
}

func (s *Student) Remove(id int) {
	for index, value := range StuList {
		if value.ID == id {
			StuList = append(StuList[:index], StuList[index+1:]...)
		}
	}
}

func (s *Student) Modefy(id int, stu Student) {
	for index, value := range StuList {
		if value.ID == id {
			StuList[index] = stu
		}
	}
}
func main() {

}
