package grades

func init() {
	students = []Student{
		{
			ID:        1,
			Firstname: "bob",
			Lastname:  "lis",
			Grades: []Grade{
				{
					Title: "Quiz",
					Type:  GradeQuiz,
					Score: 86,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz",
					Type:  GradeQuiz,
					Score: 87,
				},
			},
		},
	}
}
