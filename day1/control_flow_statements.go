package day1

import (
	"fmt"
)

var arr = []string{"A", "B", "C"}

func TryFor() {
	// k is key
	// v is value
	// range is the keyword used to iterate
	for k, v := range arr {
		fmt.Printf("index %d Value %s \n", k, v)
	}
}

func TryWhile() {
	// for is the while in go
	var i int
	for i < len(arr) {
		fmt.Printf("index %d Value %s \n", i, arr[i])
		i++
	}
}

func TryIfElse() {
	// Simple if
	//sum(1, 2, 5)
	//
	//sum(5, 5, 12)

	sum(1, 2, 5)

}

func sum(a, b, limit int) int {

	//if a+b < limit {
	//	return a + b
	//}

	// short hand notification
	if val := a + b; val < limit {
		fmt.Printf("Lenght greater than limit %d", val)
		return a + b
	} else {
		fmt.Printf("Lenght greater than %d", val)
	}

	return a + b
}

func TrySwitch(name string) {

	switch name {

	case "MAC":
		fmt.Printf("IN MAC")
	case "DELL":
		fmt.Printf("FOUND %s", name)
	default:
		fmt.Printf("Not Found %s", name)

	}

}

func TryDefer() {

	// defer Statement
	//defer fmt.Println("Control flow statements completed 1")
	//
	//defer fmt.Println("Control flow statements completed 2")
	//
	//defer fmt.Println("Control flow statements completed 3")

	// defer func

	//defer func() {
	//	fmt.Println("In defer func Control flow statements completed")
	//}()

	//this func can take inputs while initialising defer eg:
	//a := "In defer func Control flow statements completed"
	//defer func(a string) {
	//	fmt.Println(a)
	//}(a)

	// Multiple defer statement follows LIFO. eg:
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	//defer TrySwitch("DELL")
	TryFor()
	TryIfElse()
	TryWhile()
}

// to discuss
func TryString() {

	str := "hello"

	fmt.Println(&str)
	var str1 string
	str1 = "new"

	str = " world"

	//str1 := str

	fmt.Println(&str1)

	fmt.Println(&str)

}

func completed() {
	fmt.Println("In defer func Control flow statements completed")
}
