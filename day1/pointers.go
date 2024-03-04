package day1

import "fmt"

// This file covers topic such as pointers, array, maps, structs, slices and range

// Pointers
// Allows you to pass references to values and records within your program.
// Dereference using *
// Pass reference to the pointer via &

func exampleByValue(val int) {
	val++
}

func exampleByPointers(val *int) {
	*val++

}

func PointersExample() {

	x := 42
	fmt.Println("Value of x:", x)           // Output: Value of x: 42
	fmt.Println("Memory address of x:", &x) // Output : Memory reference of x

	val := 0

	fmt.Printf("Value before %d", val)

	exampleByValue(0)

	fmt.Printf("Value after pass by value %d \n", val)

	exampleByPointers(&val)
	fmt.Printf("Value after pass by reference %d", val)
}

type Person struct {
	Name string
	Age  int
}

func PointerIndirection() {
	p := &Person{Name: "Alice", Age: 30} // p is a pointer to a Person
	fmt.Println(p.Name)                  // Accessing Name field using pointer indirection
	fmt.Println(p.Age)                   // Another way to access fields using pointer indirection
}
