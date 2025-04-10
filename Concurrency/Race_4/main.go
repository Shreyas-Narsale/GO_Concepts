package main

import (
	"fmt"
	"time"
)

var counter int

func IncreCounter() {
	for i := 0; i < 1000; i++ {
		counter += i
	}
}

func main() {
	go IncreCounter()
	go IncreCounter()

	time.Sleep(10 * time.Second)
	fmt.Println(counter)
}

// cmd: go run -race main.go
// Race conditions in Go occur when two or more goroutines access the same variable concurrently,
//  and at least one of the accesses is a write
//it will not cause failure , So uncertanies in results
