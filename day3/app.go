package day3

import "github.com/self/freshers-bootcamp/database"

func GetUserNameByID(db database.Database, id int) (string, error) {
	user, err := db.GetUserByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
