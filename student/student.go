package student

import "golang.org/x/net/context"

type (
	schoolStudent interface {
		getName() string
		getSubjects() []*Subject
		getTotalScore() (float32, float32)
	}

	entityStudent struct {
		Name     string
		Subjects []*Subject
	}

	GRPCStudent struct {
		schoolStudent
	}
)

func (s entityStudent) getName() string {
	return s.Name
}

func (s entityStudent) getSubjects() []*Subject {
	return s.Subjects
}

func (s entityStudent) getTotalScore() (grandMarksTotal, grandTotal float32) {
	for _, subj := range s.Subjects {
		grandMarksTotal += subj.Marks
		grandTotal += subj.Total
	}
	return grandMarksTotal, grandTotal
}

func (stud GRPCStudent) Profile(ctx context.Context, req *ReqProfile) (res *ResProfile, err error) {
	var ok bool
	stud.schoolStudent, ok = samplesStudent[req.Id]
	res = &ResProfile{}
	if !ok {
		return res, nil
	}
	res.StudentName = stud.schoolStudent.getName()
	res.Subjects = stud.schoolStudent.getSubjects()
	res.GrandTotalMarks, res.GrandTotal = stud.schoolStudent.getTotalScore()
	return res, nil
}
