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
