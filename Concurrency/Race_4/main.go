package main

// cmd : go run --race . (race condtion check cmd)
import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var res = []int{0}

	var wg sync.WaitGroup

	var mut sync.RWMutex
	wg.Add(3)

	go func(w *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("Race 1")
		mu.Lock()
		res = append(res, 1)
		mu.Unlock()
		wg.Done()
	}(&wg, &mut)

	go func(w *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("Race 2")
		mu.Lock()
		res = append(res, 2)
		mu.Unlock()
		wg.Done()
	}(&wg, &mut)

	go func(w *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("Race 3")
		time.Sleep(5 * time.Second)
		mu.RLock()
		fmt.Println(res)
		mu.RUnlock()
		wg.Done()
	}(&wg, &mut)

	wg.Wait()
}
