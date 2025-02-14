package main

import (
	"fmt"
	"sync"
	"time"
)

func Demo(ID int, Jobs <-chan int, Result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range Jobs {
		fmt.Printf("Accepted ID : %d, Value : %d \n", ID, job)
		time.Sleep(1 * time.Second)
		Result <- job * job
		fmt.Printf("Output ID : %d, Value : %d \n", ID, job)
	}
}
func main() {
	BufferSize := 5
	WorkerSize := 3

	InputJobs := make(chan int, BufferSize)
	OutputResults := make(chan int, BufferSize)
	var wg sync.WaitGroup
	for i := 0; i < WorkerSize; i++ {
		wg.Add(1)
		go Demo(i, InputJobs, OutputResults, &wg)
	}

	for i := 0; i < BufferSize; i++ {
		InputJobs <- i
	}
	close(InputJobs)

	wg.Wait()
	close(OutputResults)

	for result := range OutputResults {
		fmt.Println("Result val:", result)
	}

}
