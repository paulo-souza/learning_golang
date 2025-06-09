package example

import "fmt"

/*
	In Golang, constant and function names have a convention for
	writing and accessing them. Names are written in camelCase or
	PascalCase. The first makes a function or variable
	inaccessible (or private) by external code, and the
	second makes them public.
*/

// Declaration and assignment specifying a type
func FirstForm() {
	const pi float32 = 3.1415
	const eightBits byte = 8
	fmt.Printf("pi = %.4f\n1 byte = %d bits\n", pi, eightBits)
}

func SecondForm() {
	const pi, eulerNumber float32 = 3.1415, 2.71828
	fmt.Printf("pi = %.4f\ne = %f\n", pi, eulerNumber)
}

func ThirdForm() {
	const (
		eightBits byte    = 8
		oneKbyte  int     = 1024
		pi        float32 = 3.1415
	)
	fmt.Printf("1 byte = %d bits\n1 Kb = %d bytes\npi = %.4f\n", eightBits, oneKbyte, pi)
}
