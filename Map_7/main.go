package main

import "fmt"

func main() {

	FavData := make(map[string]string)

	FavData["color"] = "red"
	FavData["fruit"] = "mango"
	FavData["sport"] = "cricket"

	fmt.Println(FavData)

	delete(FavData, "color")

	fmt.Println(FavData)

	for key, value := range FavData {
		fmt.Println("key:", key, ", value:", value)
	}

}
