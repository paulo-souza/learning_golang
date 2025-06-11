package example

import "fmt"

var beatles = []string{"Paul", "Ringo", "George", "John"}

/*
"Low" and "High" is a way of extracting a slice from a slice or matrix while maintaining a link to them.
"Low" is the first index of the slice and "High" is the last (but "High" does not include the index itself).
*/
func ExampleOfLowAndHigh() {
	paulAndRingo := beatles[0:2]
	georgeAndJohn := beatles[2:4]

	fmt.Println(paulAndRingo)
	fmt.Println(georgeAndJohn)

	viewBeatles := beatles[:]
	// or
	allBeatles := beatles[0:4]

	fmt.Println(viewBeatles)
	fmt.Println(allBeatles)
}

// Be careful when using "low" and "high".
func WarningWithLowAndHigh() {
	allBeatles := beatles[:]
	fmt.Println("beatles =>", beatles)
	fmt.Println("allBeatles =>", allBeatles)

	beatles[0] = "Elvis"
	beatles[2] = "Madonna"

	fmt.Println("\nbeatles =>", beatles)
	fmt.Println("allBeatles =>", allBeatles)
}
