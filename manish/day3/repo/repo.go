package repo

import (
	"github.com/jinzhu/gorm"
	"manish/day3/models"
)

func CreateStudent(student models.Student, db *gorm.DB) error {
	err := db.Table(student.TableName()).Create(&student).Error
	return err
}

func UpdateStudent(student models.Student, db *gorm.DB) error {
	err := db.Table(student.TableName()).
		Where("id = ?", student.ID).
		Updates(&student).Error
	return err
}

func FindStudent(student *models.Student, db *gorm.DB) error {
	err := db.Table(student.TableName()).
		Where("id = ?", student.ID).
		First(&student).Error
	return err
}

func UpsertStudent(student *models.Student, db *gorm.DB) error {
	err := db.Table(student.TableName()).
		Where("first_name = ?", student.FirstName).
		Where("last_name = ?", student.LastName).
		Assign(student).
		FirstOrCreate(&student).Error
	return err
}

func UpsertSubject(subject *models.Subject, db *gorm.DB) error {
	err := db.Table(subject.TableName()).
		Where("name = ?", subject.Name).
		Assign(subject).
		FirstOrCreate(&subject).Error
	return err
}

func UpsertSubjectMarks(subjectMarks *models.StudentMarks, db *gorm.DB) error {
	err := db.Table(subjectMarks.TableName()).
		Where("student_id = ?", subjectMarks.StudentId).
		Where("subject_id = ?", subjectMarks.SubjectId).
		Assign(subjectMarks).
		Updates(&subjectMarks).Error
	return err
}
