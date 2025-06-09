package example

import "fmt"

func FirstForm() {
	const sliceSize int = 5
	var sliceExample []int = make([]int, sliceSize)
	sliceExample[3] = 10
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
