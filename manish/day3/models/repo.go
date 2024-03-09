package models

import "manish/day3/config"

func CreateStudent(student *Student) error {
	if err := config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}
