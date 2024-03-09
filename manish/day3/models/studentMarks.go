package models

type StudentMarks struct {
	ID        int `json:"id"`
	StudentId int `json:"student_id"`
	SubjectId int `json:"subject_id"`
	Marks     int `json:"marks"`
}

func (s *StudentMarks) TableName() string {
	return "student_marks"
}
