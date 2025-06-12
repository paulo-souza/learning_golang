package example

import "fmt"

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
