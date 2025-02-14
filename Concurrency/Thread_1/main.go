package main

import (
	"fmt"
	"time"
)

func main() {
	go Greater("Hello")
	Greater("World")

	time.Sleep(10 * time.Second)
}

func Greater(Data string) {
	for i := 0; i < 5; i++ {
		fmt.Println(Data)
	}
}
