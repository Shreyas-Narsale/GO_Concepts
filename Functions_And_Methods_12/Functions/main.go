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

	Genric("sdbsdbj", "sbbsaj")
	Vardic(10, "sdhha", "shha")
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


func Genric[T any](inputs ...T) {
	for _, dt := range inputs {
		fmt.Println(dt)
	}
}

func Vardic(no int, strArr ...string) {
	fmt.Println("no", no, "strArr", strArr)
}
