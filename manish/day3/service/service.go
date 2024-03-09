package service

import (
	"fmt"
	"manish/day3/config"
	"manish/day3/models"
	"manish/day3/repo"
)

func CreateStudent(student models.Student) error {
	err := repo.CreateStudent(student, config.DB)
	if err != nil {
		fmt.Println("Error while creating student: ", err)
		return err
	}
	return nil
}

func UpdateStudent(student models.Student) error {
	err := repo.UpdateStudent(student, config.DB)
	if err != nil {
		fmt.Println("Error while updating student: ", err)
		return err
	}
	return nil
}

func InsertMarks(insertMarksRequest models.InsertMarksRequest) (*models.Student, error) {
	var student models.Student
	student.FirstName = insertMarksRequest.FirstName
	student.LastName = insertMarksRequest.LastName
	student.DateOfBirth = insertMarksRequest.DateOfBirth
	student.Address = insertMarksRequest.Address

	var subject models.Subject
	subject.Name = insertMarksRequest.Subject

	//Find or create student
	if err := repo.UpsertStudent(&student, config.DB); err != nil {
		fmt.Println("Error while upserting student: ", err)
		return nil, err
	}

	//Find or create subject
	if err := repo.UpsertSubject(&subject, config.DB); err != nil {
		fmt.Println("Error while upserting student: ", err)
		return nil, err
	}

	var studentMarks models.StudentMarks
	studentMarks.Marks = insertMarksRequest.Marks
	studentMarks.StudentId = student.ID
	studentMarks.SubjectId = subject.ID

	if err := repo.UpsertSubjectMarks(&studentMarks, config.DB); err != nil {
		fmt.Println("Error while upserting student_marks: ", err)
		return nil, err
	}

	fmt.Println("Student: ", student)
	fmt.Println("Subject: ", subject)
	fmt.Println("StudentMarks: ", studentMarks)

	return &student, nil
}
