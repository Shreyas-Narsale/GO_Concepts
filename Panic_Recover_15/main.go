package main

import (
	"fmt"
	"log"
)

func main() {
	defer handlePanic() // Defer the panic handling function

	// Application logic
	fmt.Println("Starting the application...")

	// Simulating a function that panics
	someFunction()

	fmt.Println("Application exited normally.")
}

func handlePanic() {
	if r := recover(); r != nil {
		// Handle the panic gracefully
		log.Printf("Recovered from panic: %v\n", r)
	}
}

func someFunction() {
	fmt.Println("Executing someFunction...")

	// Simulate a panic
	panic("Something went wrong!")
}
