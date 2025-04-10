package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter age ")
	var age int
	var name string
	fmt.Scanf("%d\n", &age)
	fmt.Println(age)

	fmt.Println("Please enter name")
	fmt.Scanf("%s\n", &name)
	fmt.Println(name)

	var status bool
	fmt.Print("Enter status (true/false): ")
	_, err := fmt.Scanf("%t\n", &status)
	if err != nil {
		fmt.Println("Invalid input! Please enter only 'true' or 'false'.")
		return
	}

	if status {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}

}
