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

    // wait till specfic time
    ch := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- "done"
    }()

    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)

    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }

    // infinite waiting loop
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    for {
        select {
        case v, ok := <-ch:
            if !ok {
                fmt.Println("Channel closed")
                return
            }
            fmt.Println(v)
        }
    }

    // if rcv not read , forgot send
    ch := make(chan int)

    select {
        case ch <- 100:
            fmt.Println("Sent 100")
        default:
            fmt.Println("Channel not ready for send")
    }

    // based on ctx for apis
    select {
        case <-ctx.Done():
            fmt.Println("Cancelled")
        case <-time.After(1 * time.Second):
            fmt.Println("Timeout!")
    }
}
