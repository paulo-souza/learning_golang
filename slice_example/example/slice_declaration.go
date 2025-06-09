package example

import "fmt"

func FirstForm() {
	const sliceSize int = 5
	var sliceExample []int = make([]int, sliceSize)
	sliceExample[3] = 10

	fmt.Println(sliceExample)
}
