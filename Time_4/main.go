package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	formatedTime := time.Now().Format("08-02-2006 Tuesday")
	fmt.Println(formatedTime)

	createdTime := time.Date(2025, time.December, 22, 12, 45, 32, 34, time.UTC)
	fmt.Println(createdTime.Format("08-02-2006 Monday"))
}
