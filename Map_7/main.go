package main

import (
	"fmt"
	"sync"
)

func main() {
	mapDt := make(map[string]string) //allowed
	//mapDt := make(map[string]int, 5) //allowed

	mapDt["key1"] = "hello"
	mapDt["key2"] = "demo"
	mapDt["key3"] = "from"

	delete(mapDt, "key2")

	for key, value := range mapDt {
		fmt.Println("key:", key, ",Value:", value)
	}

	litMapDt := map[string]int{"a": 1, "b": 2} //allowed
	fmt.Println("data[a]:", litMapDt["a"])

	//nilMap := map[string]string // not allowed

	var syncMap = sync.Map{}
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", 10)

	val, isContain := syncMap.Load("key1")
	if isContain {
		fmt.Println("key1:", val)
	}

	for key, value := range syncMap.Range {
		fmt.Println("key:", key)
		fmt.Println("value:", value)

	}
}
