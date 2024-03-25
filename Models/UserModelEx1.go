package Models

type UserEx1 struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *UserEx1) TableName() string {
	return "user_ex1"
}
