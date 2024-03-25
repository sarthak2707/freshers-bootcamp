package Models

type TestScore struct {
	ID        uint   `json:"id"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	StudentID uint   `json:"student_id" gorm:"foreignKey:StudentID; uniqueIndex:idx_student_subject"`
}

func (b *TestScore) TableName() string {
	return "test_scores"
}

func (TestScore) UniqueIndex() [][]string {
	return [][]string{
		{"StudentID", "Subject"},
	}
}
