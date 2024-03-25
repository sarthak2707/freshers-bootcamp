package Models

type Student struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	DateOfBirth int64  `json:"date_of_birth" gorm:"type:bigint"`
	Address     string `json:"address"`
}

func (b *Student) TableName() string {
	return "students"
}
