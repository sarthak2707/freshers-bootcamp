package database

type Database interface {
	GetUserByID(id int) (User, error)
}

type User struct {
	ID   int
	Name string
}

type Db struct {
}

func (d *Db) GetUserByID(id int) (User, error) {
	return User{}, nil
}
