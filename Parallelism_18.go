// Parallelism:
/* GOMAXPROCS: Used to achieve parallelism
	It defines the maximum number of CPU cores that a Go program can utilize at the same time
	Each core executes one goroutine at a time, so if there are 8 cores, up to 8 goroutines can run simultaneously
*/
/*
Concurrency (1 CPU)
	dealing with many tasks at once
	Task A → Task B → Task A → Task C → Task B
	Switching (context switching)
Parallelism (multi CPU)
	doing many tasks at the exact same time
	CPU1 → Task A
	CPU2 → Task B
	CPU3 → Task C
*/

package main

import (
	"fmt"
	"runtime"
)

func work(id int) {
	fmt.Println("Working:", id)
}

func main() {
	runtime.GOMAXPROCS(4)

	for i := 0; i < 1000; i++ {
		go work(i)
	}

	select {}
}


// Simple Visualization
// Case 1:
// runtime.GOMAXPROCS(1)
// 1000 goroutines
//    ↓
// Only 1 runs at a time
//    ↓
// Others wait (fast switching)

// 👉 Concurrency = 1000
// 👉 Parallelism = 1

// Case 2:
// runtime.GOMAXPROCS(4)
// 1000 goroutines
//    ↓
// 4 run simultaneously
//    ↓
// 996 waiting

// 👉 Concurrency = 1000
// 👉 Parallelism = 4
