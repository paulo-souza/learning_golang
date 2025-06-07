package example

import "fmt"

/*
	In Golang, constant and function names have a convention for
	writing and accessing them. Names are written in camelCase or
	PascalCase. The first makes a function or variable
	inaccessible (or private) by external code, and the
	second makes them public.
*/

func FirstForm() {
	const pi float32 = 3.1415
	const eightBits byte = 8
	fmt.Printf("pi = %.4f\n1 byte = %d bits\n", pi, eightBits)
}
