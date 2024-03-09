package models

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Subject) TableName() string {
	return "subjects"
}
