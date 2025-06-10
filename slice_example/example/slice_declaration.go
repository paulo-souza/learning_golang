package example

import "fmt"

// Example of how to declare a slice specifying a type, size and automatically getting a default value according to its type.
func FirstForm() {
	const sliceSize int = 5
	var sliceExample []int = make([]int, sliceSize)
	fmt.Println(sliceExample)
}

// Example of how to briefly declare a slice by specifying a size and automatically getting a default value relative to the type.
func SecondForm() {
	const sliceSize int = 6
	sliceExample := make([]int, sliceSize)
	fmt.Println(sliceExample)
}

// Example of how to create a slice in a brief way already initialized with elements.
func ThirdForm() {
	sliceExample := []int{0, 0, 0, 10, 0}
	fmt.Println(sliceExample)
}

// Example of how to create a slice through a link with an array.
func FourthForm() {
	anArrayOfIntegers := [5]int{0, 0, 0, 10, 0}
	sliceExample := anArrayOfIntegers[0:5]
	//or
	// sliceExample := anArrayOfIntegers[:]
	//or
	// sliceExample := anArrayOfIntegers[:5]
	//or
	// sliceExample := anArrayOfIntegers[0:]

	fmt.Println(sliceExample)
}

// Declaration-only example, which automatically gets an empty slice as the default assigned value.
func FifthForm() {
	var sliceExample1 []int
	fmt.Println(sliceExample1)

	var sliceExample2 []int = make([]int, 0)
	fmt.Println(sliceExample2)

	sliceExample3 := []int{}
	fmt.Println(sliceExample3)
}

// Sample assignment of an element to an existing slice.
func Samples() {
	var sliceExample1 []int = make([]int, 4)
	sliceExample1[1] = 7
	fmt.Println(sliceExample1)

	sliceExample2 := []int{0, 0, 0, 0}
	sliceExample2[2] = 10
	fmt.Println(sliceExample2)
}
