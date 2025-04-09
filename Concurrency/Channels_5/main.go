package main

import (
	"fmt"
	"sync"
)

func main() {

	//channel := make(chan int) //unbuffered Channel
	channel := make(chan int, 2) //buffered Channel

	var wg sync.WaitGroup
	wg.Add(2)

	//RCV Channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-ch
		fmt.Println("isChanOpen:", isChanOpen)
		fmt.Println("Val:", val)
		defer wg.Done()
	}(channel, &wg)

	//SEND Channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		close(ch)
		// ch <- 6
		defer wg.Done()
	}(channel, &wg)

	wg.Wait()

	//
	// cha := make(chan int)
	// fmt.Println(<-cha)
	// cha <- 5

	// Deadlock Issue : channel is unbuffered, so need a receiver to be ready before sending.

	channel = make(chan int)
	go func() {
		channel <- 5 // Send in a goroutine
	}()
	fmt.Println(<-channel) // Receive and print
	//Allowed As receiver is ready

	cha1 := make(chan int, 1)
	cha1 <- 5
	fmt.Println(<-cha1)
	//Allowed bcuz it store data in buffer
}
