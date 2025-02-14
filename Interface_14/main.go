package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Paramter() float64
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return ((s.Length) * (s.Length))
}

func (s Square) Paramter() float64 {
	return 4 * s.Length
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Paramter() float64 {
	return 2 * r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Paramter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var S Shape

	S = Circle{10}
	fmt.Println(S.Area())
	fmt.Println(S.Paramter())

	S = Square{10}
	fmt.Println(S.Area())
	fmt.Println(S.Paramter())

	S = Rectangle{10, 20}
	fmt.Println(S.Area())
	fmt.Println(S.Paramter())

}
