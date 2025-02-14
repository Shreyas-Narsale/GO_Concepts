package main

import "fmt"

func main() {

	var days = []string{"mon", "tue", "wed", "thurs", "fri", "sat", "sun"}

	//for using index
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for each using index
	for i := range days {
		fmt.Println(days[i])
	}

	//for each
	for index, day := range days {
		fmt.Println("index:", index, ",day :", day)
	}

	// while (these is no while loop directly but we can implement using for)

	j := 0
	for j < len(days) {
		fmt.Println(days[j])
		j++
	}

	fmt.Println("Testing goto , continue and break ")
	j = 0
	for j < len(days) {
		// if j == 2 {
		// 	j++
		// 	continue
		// }
		// if j == 6 {
		// 	break
		// }
		if j == 4 {
			goto loc
		}
		fmt.Println(days[j])
		j++
	}
loc:
	fmt.Println("Inside label loc")

	// Do while  (these is no do while loop directly but we can implement using for)

	m := 0

	for {
		fmt.Println(m)
		m++

		if m == 10 {
			break
		}
	}
}
