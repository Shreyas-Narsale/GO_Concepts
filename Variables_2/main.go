package main

import (
	"fmt"
)

// SerailNo := 123 // Not allowed
var serailNo int = 123 //private Accessible only within this package

const Token string = "csdjbsdnkjwd" //public Accessible in Other package

func main() {
	var username string = "Hello"
	fmt.Println(username)
	fmt.Printf("is Type of %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("is Type of %T\n", isLoggedIn)

	// var iVal uint8 = 256 // Not allowed as uint8 unsigned int of 8 bit so range 0 to 255
	var iVal uint8 = 255
	fmt.Println(iVal)
	fmt.Printf("is Type of %T\n", iVal)

	var smallFloat float32 = 155.993933993484999393993948484833939384848393
	fmt.Println(smallFloat) // output 155.99393
	fmt.Printf("is Type of %T\n", smallFloat)

	var largeFloat float64 = 155.993933993484999393993948484833939384848393
	fmt.Println(largeFloat) // output 155.993933993485
	fmt.Printf("is Type of %T\n", largeFloat)

	//default values
	var iNo int
	fmt.Println(iNo) //0

	//implicit type
	//short variable declaration operator
	iNum := 123
	fmt.Println(iNum)

	serailNo = 456
	fmt.Println(serailNo)

	const value = 123 // const variable
	fmt.Println(value)

	// Token = "djjsj" not allowed const variable
	fmt.Println(Token)

	var ch rune //size 4 Bytes
	ch = 'A'
	fmt.Printf("%c", ch)

	var byt byte //size 1 Bytes
	byt = 'A'
	fmt.Printf("%c", byt)

	ch = 'म' // as run is of 4 bytes so allowed
	fmt.Printf("%c", ch)

	// byt = 'म' not allowed multi byte unicode , So required memory of 4 bytes
	//fmt.Printf("%c", byt)
}
