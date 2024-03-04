package day1

import "fmt"

// Maps
// Maps are key value pairs or associated data type or dict(in other languages)
func MapExample() {
	// Create an empty map with string keys and int values
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["orange"] = 30

	// Access values from the map
	fmt.Println("Value of apple:", myMap["apple"])
	fmt.Println("Value of banana:", myMap["banana"])
	fmt.Println("Value of orange:", myMap["orange"])

	// Update the value of a key
	myMap["banana"] = 25
	fmt.Println("Updated value of banana:", myMap["banana"])

	// Check if a key exists in the map
	_, ok := myMap["apple"]
	fmt.Println("Is apple present in the map?", ok)

	// Delete a key from the map
	delete(myMap, "orange")
	fmt.Println("Map after deleting orange:", myMap)

}

// Define another struct type named `Address`
type Address struct {
	Street string
	City   string
}

func StructExample() {
	person1 := Person{
		Name: "Alice",
		Age:  30,
	}

	// Create a struct instance using the `new` function (returns a pointer)
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 25

	person3 := Person{"Charlie", 35}

	// Create a pointer to a struct using the `&` operator
	person4 := &Person{
		Name: "David",
		Age:  40,
	}

	// Print the struct instances
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", *person2) // Dereference the pointer to access the struct
	fmt.Println("Person 3:", person3)
	fmt.Println("Person 4:", *person4) // Dereference the pointer to access the struct
}

func RangeExample() {

	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["orange"] = 30

	for k, v := range myMap {
		fmt.Printf("key %s value %d", k, v)
	}

}
