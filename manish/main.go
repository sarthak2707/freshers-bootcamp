package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"manish/day3/config"
	"manish/day3/models"
	"manish/day3/routes"
)

var err error

func day3App() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Error while initializing DB: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Student{})
	config.DB.AutoMigrate(&models.Subject{})
	config.DB.AutoMigrate(&models.StudentMarks{})

	r := routes.SetupRouter()
	r.Run()
}

func main() {
	//question 1
	//day1.Question1()

	//question 2
	//day1.Question2()

	//question 3
	//day1.Question3()

	// day2 Ques1
	//day2.Question1()

	//day2.Question1_2()

	// day2 Ques2
	//day2.Question2()

	//day2 Ques3
	//day2.Question3()

	//day2.ChannelWithDeadlockExample2()

	day3App()

}
