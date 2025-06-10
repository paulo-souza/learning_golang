package example

import "fmt"

// Example of how to declare a slice specifying a type, size and automatically getting a default value according to its type.
func FirstForm() {
	const sliceSize int = 5
	var sliceExample []int = make([]int, sliceSize)
	fmt.Println(sliceExample)
}

func SecondForm() {
	const sliceSize int = 6
	sliceExample := make([]int, sliceSize)
	sliceExample[4] = 9
	fmt.Println(sliceExample)
}

func ThirdForm() {
	sliceExample := []int{0, 0, 0, 10, 0}
	fmt.Println(sliceExample)
}

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

func FifthForm() {
	var sliceExample1 []int
	fmt.Println(sliceExample1)

	var sliceExample2 []int = make([]int, 0)
	fmt.Println(sliceExample2)

	sliceExample3 := []int{}
	fmt.Println(sliceExample3)
}
