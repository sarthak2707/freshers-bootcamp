package day1

import (
	"errors"
	"fmt"
)

func defaultErrorExample(val bool) (int, error) {
	if val {
		return 0, errors.New("default error returned")
	}
	return 1, nil
}

func ErrorExample() {
	val, err := defaultErrorExample(false)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

}

// CustomError is a custom error type
type CustomError struct {
	message string
}

// CustomError is a custom error type
type CustomError2 struct {
	message string
}

// Error method implements the error interface for CustomError
func (ce CustomError) Error() string {
	return ce.message
}

// Error method implements the error interface for CustomError
func (ce CustomError2) Error() string {
	return ce.message
}

// Function that returns a custom error
func doSomethingCustomError(val bool) error {
	if val {
		return CustomError{"Something went wrong custom error 1"}
	}

	return CustomError2{"Something went wrong custom error 2"}

}

func CustomErrorExample() {
	err := doSomethingCustomError(false)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
