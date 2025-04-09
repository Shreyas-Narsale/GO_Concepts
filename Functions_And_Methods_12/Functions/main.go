package main

import "fmt"

func main() {
	fmt.Println("Inside main")
	Demo()

	func() {
		fmt.Println("Inside main another function")
	}()

	iRes := Sum(10, 20)
	fmt.Println(iRes)

	iRes = Add(2, 3, 4, 5, 6)
	fmt.Println(iRes)
}

func Demo() {
	fmt.Println("Inside Demo")
}

func Sum(value1, value2 int) int {
	fmt.Println("Inside Sum")
	return value1 + value2
}

//varadic Function
func Add(values ...int) int {
	res := 0
	fmt.Println("Inside Add")
	for _, val := range values {
		res = res + val
	}
	return res
}
