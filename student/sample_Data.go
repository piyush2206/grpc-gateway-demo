package student

var (
	samplesStudent = map[string]*entityStudent{
		"0001": &entityStudent{
			Name: "Piyush",
			Subjects: []*Subject{
				{
					SubjectName: "Maths",
					Total:       150,
					Marks:       125,
				},
				{
					SubjectName: "Science",
					Total:       100,
					Marks:       67,
				},
			},
		},
		"0002": &entityStudent{
			Name: "Abhinav",
			Subjects: []*Subject{
				{
					SubjectName: "Maths",
					Total:       150,
					Marks:       140,
				},
				{
					SubjectName: "Science",
					Total:       100,
					Marks:       88,
				},
			},
		},
	}
)
