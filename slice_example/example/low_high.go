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
	fmt.Println("\nOBS.: Para solucionar este problema pode se fazer uso de \"copy\" ou \"append\"")
}

// Example of using append and one of the solutions to problems that "low" and "high" can cause.
func UsingAppend() {
	newBeatles := append(beatles, "Elvis", "Madonna")
	beatles[2] = "Michael"
	fmt.Println("beatles =>", beatles)
	fmt.Println("newBeatles =>", newBeatles)
}

// Digging deeper into the use of append and demonstrating interesting things about it.
func OtherInterestingThingsAboutAppend() {
	sliceExample := make([]int, 3, 5)
	sliceExample[0] = 21
	sliceExample[1] = 5
	sliceExample[2] = 9
	fmt.Println("sliceExample =", sliceExample)
	fmt.Println("Tamanho de sliceExample:", len(sliceExample))
	fmt.Println("Capacidade de sliceExample:", cap(sliceExample))

	//Adds elements 90 and 77 in "sliceExample" and only changes its size
	sliceExample = append(sliceExample, 90, 77)
	fmt.Println("\nsliceExample =", sliceExample)
	fmt.Println("Tamanho de sliceExample:", len(sliceExample))
	fmt.Println("Capacidade de sliceExample:", cap(sliceExample))

	//Creates and assigns a new slice in "sliceExample" with new size and new capacity
	sliceExample = append(sliceExample, 13, 43, 69)
	fmt.Println("\nsliceExample =", sliceExample)
	fmt.Println("Tamanho de sliceExample:", len(sliceExample))
	fmt.Println("Capacidade de sliceExample:", cap(sliceExample))

	/*
		Copies the elements of "sliceExample" to a new slice and adds more elements to the end of it,
		such as 90, 77, 3 and 1, and then assigns them to the variable "otherSlice"
	*/
	otherSlice := append(sliceExample, 90, 77, 3, 1)
	fmt.Println("\notherSlice =", otherSlice)
	fmt.Println("Tamanho da otherSlice:", len(otherSlice))
	fmt.Println("Capacidade da otherSlice:", cap(otherSlice))
}
