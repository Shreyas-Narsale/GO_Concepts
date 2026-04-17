package main

import (
    "fmt"
    "time"
)

func main() {
    // select : wait on multiple channel at the same time
    // Wait until any one of several channel operations is ready, then execute that case.
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() { ch1 <- 1 }()
    go func() { ch2 <- 2 }()
    
    select {
    case v := <-ch1:
        fmt.Println("From ch1:", v)
    
    case v := <-ch2:
        fmt.Println("From ch2:", v)
    }
    // Non-blocking select : default excutes if no one is read 

    ch := make(chan int)
    select {
    case v := <-ch:
        fmt.Println(v)
    default:
        fmt.Println("No data available")
    }
}
