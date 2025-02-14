package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello , Jay Ganesh"
	fileName := "./log.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	len, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No. of Bytes writtern", len)
	ReadFile(fileName)
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Unable to read file")
	}

	fmt.Println("Data is", string(data))
}
