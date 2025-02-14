package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Give rating:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	// comma ok || err
	fmt.Println("Input Data is :", input)
	fmt.Printf("Type of input %T \n", input)

	//Conversion
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("error occured", err)
	}
	fmt.Println("User has given rating ", val, " *")

}
