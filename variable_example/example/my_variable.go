package example

import "fmt"

// Just declaring
func FirstForm() {
	var oneInt int
	var oneflot32 float32
	var oneflot64 float64
	var aString string

	fmt.Printf("oneInt = %v\n", oneInt)
	fmt.Printf("oneflot32 = %v\n", oneflot32)
	fmt.Printf("oneflot64 = %v\n", oneflot64)
	fmt.Printf("aString = %v\n", aString)
}

// Typed declaration and assignment
func SecondForm() {
	var numberTwo int = 2
	var numberFloat32 float32 = 3.334
	var numberFloat64 float64 = 3.1415
	var otherString string = "Anything string"

	fmt.Printf("numberTwo = %d\n", numberTwo)
	fmt.Printf("numberFloat32 = %.3f\n", numberFloat32)
	fmt.Printf("numberFloat64 = %.4f\n", numberFloat64)
	fmt.Printf("otherSting = %s\n", otherString)

}

// Syntactic sugar for declaration and assignment
func ThirdForm() {
	numberFive := 5
	otherNumberFloat32 := float32(1.5967)
	otherNumberFloat64 := 0.53
	anyString := "any string"

	fmt.Printf("numberFive = %d\n", numberFive)
	fmt.Printf("otherNumberFloat32 = %.4f\n", otherNumberFloat32)
	fmt.Printf("otherNumberFloat64 = %.2f\n", otherNumberFloat64)
	fmt.Printf("anyString = %s\n", anyString)
}

// Declare in a grouped manner
func FourthForm() {
	var (
		anyString       string  = "any string"
		someNumber              = 10
		someFloatNumber float64 = 3.1415
	)

	fmt.Printf("anyString = %s\n", anyString)
	fmt.Printf("someNumber = %d\n", someNumber)
	fmt.Printf("someFloatNumber = %.4f\n", someFloatNumber)

}

// Multiple inline statement
func FifthForm() {
	var anyString, someString string
	var anyNumber, someNumber int

	fmt.Printf("anyString = %s\n", anyString)
	fmt.Printf("someString = %s\n", someString)
	fmt.Printf("anyNumber = %d\n", anyNumber)
	fmt.Printf("someNumber = %d\n", someNumber)
}

// Multiple declaration and assignment inline
func SixthForm() {
	var anyString, someString string = "any string", "some string"
	fmt.Println("anyString =", anyString)
	fmt.Println("someString =", someString)
}

// Multiple declaration and assignment with type inference
func SeventhForm() {
	var anyFloatNumber, someFloatNumber = 3.14159, 2.7182
	fmt.Printf("anyFloatNumber = %.5f\n", anyFloatNumber)
	fmt.Printf("someFloatNumber = %.5f\n", someFloatNumber)
}

// Syntactic sugar for declaring and assigning multiple values
func EighthForm() {
	someString, someFloatNumber := "some string", 3.1415
	fmt.Printf("someString = %s\n", someString)
	fmt.Printf("someFloatNumber = %.4f\n", someFloatNumber)
}

func OtherSamples() {
	stringWithLineBreak := `primeira linha
	segunda linha com quebra
	terceira linha com quebra 
	quarta linha com quebra
	`
	fmt.Print(stringWithLineBreak)
}

//TODO: Refatorar para métodos como primeiraForma, segundaForma, terceiraForma
