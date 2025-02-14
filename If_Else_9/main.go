package main

import "fmt"

func main() {
	var age = 10

	if age < 18 {
		fmt.Println("Under age")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	if Num := 3; Num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
