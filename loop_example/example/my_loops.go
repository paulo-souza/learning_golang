package example

import "fmt"

/*
	Every repetition structure in Golang is defined with the reserved word "for",
	and there are several ways to do this, below are the main ones.
*/

func FirstForm() {
	count := 1
	for count <= 10 {
		fmt.Println(count)
		count = count + 1
	}
}

func SecondForm() {
	for count := 1; count <= 10; count++ {
		fmt.Println(count)
	}
}

func ThirdForm() {
	var myArray [10]int = [10]int{5, 6, 3, 1, 9, 2, 8, 5, 4, 10}

	for index, value := range myArray {
		fmt.Printf("myArray[%d] = %d\n", index, value)
	}
}
