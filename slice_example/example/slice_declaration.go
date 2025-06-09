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
