package main

import "fmt"

func main() {
	Day := "Wednesday"

	switch Day {
	case "Monday":
		fmt.Println("it's Monday")
		//break  (not required in golang it's bydefault use break)
	case "Tuesday":
		fmt.Println("it's Tuesday")
	case "Wednesday":
		fmt.Println("it's Wednesday")
		fallthrough //if we need to continue to next case
	case "Thursday":
		fmt.Println("it's Thursday")
	case "Friday":
		fmt.Println("it's Friday")
	case "Saturday":
		fmt.Println("it's Saturday")
	case "Sunday":
		fmt.Println("it's Sunday")
	default:
		fmt.Println("Please provide correct day")
	}
}
