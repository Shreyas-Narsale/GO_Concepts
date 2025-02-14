package main

import "fmt"

func main() {
	//similar like c
	var ptr *int
	fmt.Println("Data of ptr", ptr)
	//fmt.Println("Values Inside ptr", *ptr) gives error as not assigned any values  (invalid memory address)

	iNo := 10
	ptr = &iNo
	fmt.Println("Data of ptr", ptr)
	fmt.Println("Values Inside ptr", *ptr)

	*ptr = (*ptr) * 2
	fmt.Println("Values Inside ptr", *ptr)
}
