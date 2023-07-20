package grades

import "fmt"

type Student struct {
	ID        int
	Firstname string
	Lastname  string
	Grades    []Grade
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / float32(len(s.Grades))
}

type Students []Student

func (st Students) GetByID(id int) (*Student, error) {
	for i := range st {
		if st[i].ID == id {
			return &st[i], nil
		}
	}
	return nil, fmt.Errorf("Student with ID %d not found", id)
}

var students Students
