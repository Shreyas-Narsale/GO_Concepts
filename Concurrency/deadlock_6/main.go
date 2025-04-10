package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 42

	fmt.Println(<-ch)
}

//The line ch <- 42 sends data into an unbuffered channel.

//But there's no goroutine waiting to receive from that channel.

//So the sender blocks forever, and Go panics with a deadlock error.

//deadlock:
//Goroutines wait forever for a resource that never becomes free
