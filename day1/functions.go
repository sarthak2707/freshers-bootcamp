package day1

import "fmt"

//functions in go
//define with keyword func
// eg : func (r Receiver) funcName() (return parameters)
// Functions starting with capital Letter is Exported
// Such functions can be only be accessed outside the package

// Can return multiple parameters
func exampleWithMultipleReturnValues() (int, bool) {
	fmt.Println("in func")
	return 0, false
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// Not exported package. Cannot be accessed outside the package
func exampleFunction(a, b int) {

}

//func exampleFunction(a, b, c int) {
//
//}

type ExampleFunc struct {
}

// function callers are called as receivers
func (e ExampleFunc) ExampleFunctionWithReceiver() {

}

// Naked return
// return statement without arguments returns the named return values
func ExampleFunctionNamedReturn(c int, d int) (a, b int) {
	a = c * c
	b = d * d
	return
}

// Variadic Params
func ExampleFunctionWithVariadicParams(arr ...string) {
	for _, val := range arr {
		fmt.Printf("variadic params example %s \n", val)
	}
}

func ExampleFunctionClosure() func() int {
	i := 1

	return func() int {
		i = i * 2
		fmt.Printf("Next Execution %d \n", i)
		return i
	}
}

func ExampleClosure() {

	sum := func(a int, b int) {
		fmt.Println(a + b)
	}

	sum(1, 2)

	closureFunc := ExampleFunctionClosure()

	fmt.Println(closureFunc())

	fmt.Println(closureFunc())
}

func ExampleFunctionRecursion() {

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(5))
}

func FunctionExample() {
	exampleWithMultipleReturnValues()
}
