package example

import "fmt"

// Example of using copy and one of the solutions to problems that "low" and "high" can cause.
func UsingCopy() {
	var originalBeatles []string = make([]string, 4)
	copy(originalBeatles, beatles)

	beatles[0] = "Elvis"
	beatles[2] = "Madonna"

	fmt.Println("beatles =>", beatles)
	fmt.Println("originalBeatles =>", originalBeatles)
}
