package Models

import (
	"fmt"
	"freshers-bootcamp/Config"
)

// GetAllStudents Fetch all students data
func GetAllStudents(students *[]Student) (err error) {
	if err = Config.DB.Find(students).Error; err != nil {
		return err
	}
	return nil
}

// CreateStudent ... Insert New data
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// GetStudentById ... Fetch only one student by Id
func GetStudentById(student *Student, id uint) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent ... Update student
func UpdateStudent(student *Student, id uint) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}
