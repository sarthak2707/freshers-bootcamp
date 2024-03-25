package Models

type TestScore struct {
	Id        uint    `json:"id"`
	Subject   string  `json:"name"`
	Marks     int     `json:"last_name"`
	StudentId uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentId"`
}

func (b *TestScore) TableName() string {
	return "test_scores"
}
