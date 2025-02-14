package main

import (
	"fmt"
	"sort"
)

func main() {

	// Array member by member inilization

	var fruitlist [5]string

	fruitlist[0] = "mango"
	fruitlist[1] = "apple"
	fruitlist[2] = "banana"

	fmt.Printf("Type of fruitlist %T\n", fruitlist)
	fmt.Println("All items", fruitlist)
	fmt.Println("len", len(fruitlist))
	fmt.Println("Capacity", cap(fruitlist))

	// Array member by list

	var fruitlist1 = [5]string{"apple", "banana", "mango"}

	fmt.Printf("Type of fruitlist %T\n", fruitlist1)
	fmt.Println("All items", fruitlist1)
	fmt.Println("len", len(fruitlist1))
	fmt.Println("Capacity", cap(fruitlist1))

	// Slices
	// member by member inilization
	var colorlist []string

	colorlist = append(colorlist, "red")
	colorlist = append(colorlist, "yellow")
	fmt.Printf("Type of fruitlist %T\n", colorlist)
	fmt.Println("All items", colorlist)
	fmt.Println("len", len(colorlist))      // 2
	fmt.Println("Capacity", cap(colorlist)) //2

	colorlist = append(colorlist, "pink")
	colorlist = append(colorlist, "white")
	colorlist = append(colorlist, "black")
	fmt.Println("All items", colorlist)
	fmt.Println("len", len(colorlist))      //5
	fmt.Println("Capacity", cap(colorlist)) // 8 (Capacity is increased in 2nth)

	// Slice of Slice

	colorlist = append(colorlist[1:])
	fmt.Println("All items", colorlist)

	colorlist = append(colorlist[:3])
	fmt.Println("All items", colorlist)

	colorlist = append(colorlist[1:2]) //last index value is not inclued
	fmt.Println("All items", colorlist)

	// Dynmaic Slices using make

	var highScores = make([]int, 3)
	highScores[0] = 90
	highScores[1] = 20
	highScores[2] = 50

	// highScores[3] = 40, not allowed
	highScores = append(highScores, 45, 55)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// remove values in slices

	var courses = []string{"c", "cpp", "java", "ruby", "go", "html"}

	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
