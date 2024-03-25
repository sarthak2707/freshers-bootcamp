package main

import (
	"fmt"
	"freshers-bootcamp/Config"
	"freshers-bootcamp/Models"
	"freshers-bootcamp/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	//day1Exercises.MatrixExec()
	//day1Exercises.ExpressionTree()
	//day1Exercises.SalaryCalculator()
	//day2Exercises.LetterCounter()
	//day2Exercises.TeacherRating()
	//day2Exercises.BankBalanceUpdate()

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.UserEx1{})
	r := Routes.SetupRouter()
	//running
	r.Run()

	// var declaration

	//var (
	//	myVariable int
	//	name string
	//)
	//myVariable = 10
	//fmt.Print(myVariable)

	//day1.TryFor()
	//day1.TryWhile()

	//day1.TryIfElse()
	//day1.TrySwitch("jafbkjaf")

	//day1.TryDefer()

	//day1.TryString()

	//day1.FunctionExample()
	//val := []string{"1", "2"}
	//day1.ExampleFunctionWithVariadicParams("1", "2", "3")

	//a, b := day1.ExampleFunctionNamedReturn(4, 5)
	//
	//fmt.Println(a)
	//fmt.Println(b)

	//day1.ExampleClosure()

	//day1.ExampleFunctionRecursion()

	//re := day1.ExampleFunc{}
	//
	//re.ExampleFunctionWithReceiver()

	//day1.InterfaceExample()

	//day1.InterfaceValueAndType()

	//day1.EmptyInterfaceExample()

	//day1.TypeAssertionInterface()

	//day1.ExampleSlice()

	//day1.SliceResizing()

	//day1.MapExample()

	//day1.StructExample()

	//day1.RangeExample()

	//day1.PointersExample()

	//day1.PointerIndirection()

	//day1.ErrorExample()

	//day1.CustomErrorExample()

	//day2.GoRoutinesExample()

	// unbuffered
	//day2.ChannelWithSenderAndReceiver()

	//day2.BufferedChannelExample()

	//day2.ChannelSynchronization()

	//day2.ChannelsWithSelectStatement()

	//day2.NonBlockingChannels()

	//day2.CloseChannels()

	//day2.ChannelWithDeadlock()

	//day2.ChannelWithDeadlockExample2()

	//day2.RangeOverChannels()

	//day2.WaitGroup()

	//day2.AtomicInt()

	//day2.MutexWithRoutines()

	//day2.StatefulGoRoutine()

	//day2.Signals()

	//day2.WorkerPoolExample()

	//day3.ContextExample()

	//day3.InitialiseRoute()
}
