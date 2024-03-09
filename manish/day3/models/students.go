package models

type Student struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name" gorm:"unique"`
	LastName    string `json:"last_name"`
	DateOfBirth int    `json:"date_of_birth"`
	Address     string `json:"address"`
}

func (s *Student) TableName() string {
	return "students"
}
