package day1

import "fmt"

// ExampleArray Arrays have fixed size
// Cannot resize the array
func ExampleArray() {
	// Array declaration and initialization
	var arr [5]int // Declare an array of length 5 with zero values
	arr[0] = 1     // Assign value 1 to the first element
	arr[1] = 2     // Assign value 2 to the second element
	arr[2] = 3     // Assign value 3 to the third element
	arr[3] = 4     // Assign value 4 to the fourth element
	arr[4] = 5

	// Assign value 5 to the fifth element

	fmt.Println("Array:", arr) // Output: Array: [1 2 3 4 5]
}

// Slices are dynamically sized views into arrays and are more flexible
// Keeps reference to the underlying array
func ExampleSlice() {
	// Slice declaration and initialization
	slice := []int{1, 2, 3, 4, 5} // Declare and initialize a slice using a literal

	fmt.Println("Slice:", slice)         // Output: Slice: [1 2 3 4 5]
	fmt.Println("Length:", len(slice))   // Output: Length: 5
	fmt.Println("Capacity:", cap(slice)) // Output: Capacity: 5

	// append func to add values
	// need to assign append value to a variable or slice
	// why do we assingn slice again while using append
	slice = append(slice, 6)

	fmt.Println("Slice:", slice)
	fmt.Println("Cap Slice:", cap(slice))

}

// Slice Resizing
func SliceResizing() {

	// Create a slice with initial capacity of 3 and length of 0
	slice := make([]int, 0, 3)

	// Loop from 1 to 5
	for i := 1; i <= 7; i++ {
		// Append the current value of i to the slice
		slice = append(slice, i)
		fmt.Printf("Slice after appending %d: Length = %d, Capacity = %d, Elements = %v\n", i, len(slice), cap(slice), slice)
	}

	exampleWithMultipleReturnValues()

}
