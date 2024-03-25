package Models

import (
	"fmt"
	"freshers-bootcamp/Config"
)

func GetScoresByStudentId(scores *[]TestScore, id uint) (err error) {
	if err = Config.DB.Where("student_id = ?", id).Find(scores).Error; err != nil {
		return err
	}
	return nil
}

func GetScoreForStudentBySubject(score *TestScore, id uint, subject string) (err error) {
	if err = Config.DB.Where("student_id = ? and subject = ?", id, subject).First(score).Error; err != nil {
		return err
	}
	return nil
}

func CreateTestScore(score *TestScore) (err error) {
	if err = Config.DB.Create(score).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTestScore(score *TestScore) (err error) {
	fmt.Println(score)
	err = Config.DB.Save(score).Error
	return err
}
