package day1

import (
	"fmt"
	"math"
	"reflect"
)

//An interface type is defined as a set of method signatures.
// A type implements an interface by implementing its methods.
//There is no explicit declaration of intent, no "implements" keyword.
// interface values can be thought of as a tuple of a value and a concrete type

// Define an interface named `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

type empty interface {
}

// Define a struct type named `Circle`
type Circle struct {
	Radius float64
}

// Define a method for `Circle` that satisfies the `Shape` interface
func (c Circle) Area() float64 {
	fmt.Println("In circle Area")
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	fmt.Println("In circle Peri")
	return 2 * math.Pi * c.Radius
}

// Define a struct type named `Rectangle`
type Rectangle struct {
	Width  float64
	Height float64
}

// Define a method for `Rectangle` that satisfies the `Shape` interface
func (r Rectangle) Area() float64 {
	fmt.Println("In Rect Area")
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("In Rect Peri")
	return 2*r.Width + 2*r.Height
}

func InterfaceExample() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Declare a variable of type `Shape` and assign `Circle` to it
	var shape Shape
	shape = circle

	// Call methods defined by the `Shape` interface
	fmt.Printf("Area of the circle: %.2f\n", shape.Area())
	fmt.Printf("Perimeter of the circle: %.2f\n", shape.Perimeter())

	// Assign `Rectangle` to the `Shape` variable
	shape = rectangle

	// Call methods defined by the `Shape` interface
	fmt.Printf("Area of the rectangle: %.2f\n", shape.Area())
	fmt.Printf("Perimeter of the rectangle: %.2f\n", shape.Perimeter())
}

func printValue(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func EmptyInterfaceExample() {
	printValue(42)
	printValue("Hello")
	printValue(true)
}

func InterfaceValueAndType() {

	var shape Shape

	shape = Circle{}
	fmt.Println("Type:", reflect.TypeOf(shape))
	fmt.Println(shape.Area())

	shape = Rectangle{}
	fmt.Println("Type:", reflect.TypeOf(shape))
	fmt.Println(shape.Area())

}

func TypeAssertionInterface() {
	var val interface{}

	val = 42

	c, ok := val.(int)
	if ok {
		fmt.Println(c)
	}

	if v, ok := val.(int); ok {
		fmt.Println("Value is an int:", v)
	}

	val = "hello"
	if v, ok := val.(string); ok {
		fmt.Println("Value is a string:", v)
	}
}

// to discuss
func PointerReceiverWithNilInterfaceValue() {

	var ptr *Circle // Declaring a nil pointer

	// Calling method with pointer receiver on nil pointer
	ptr.Area()
}
