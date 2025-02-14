package main

import "fmt"

func main() {

	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("1")
	defer fmt.Println("World")
	MyDeffer()
	fmt.Println("Hello")

}

func MyDeffer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
